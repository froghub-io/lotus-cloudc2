package ffiwrapper

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 标识
const CloudC2 = "cloud c2"

// 任务状态
const (
	Fail    = "fail"
	Being   = "being"
	Running = "running"
	Success = "success"
	Empty   = "empty"
)

// 错误类型
const (
	WaitCommit2 = "WaitCommit2"
	ProofError  = "cloud c2 proof error"
	ProofFail   = "cloud c2 failed"
)

type c2BodyRequest struct {
	ActorID      uint64 `json:"actor_id"`
	SectorID     uint64 `json:"sector_id"`
	Phase1Out    []byte `json:"phase_1_out"`
	Phase1OutMd5 string `json:"phase_1_out_md5"`
}

type c2BodyResponse struct {
	Proof []byte `json:"proof"`
	State string `json:"state"`
}

func requestHttpWithSign(token string, url string, body interface{}) ([]byte, error) {
	data, err := signBody(body, token)
	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var zBuf bytes.Buffer
	zw := gzip.NewWriter(&zBuf)
	if _, err = zw.Write(j); err != nil {
		return nil, err
	}
	zw.Close()
	req, err := http.NewRequest(http.MethodPost, url, &zBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%v", string(respBody)))
	}
	return respBody, nil
}

type mapEntryHandler func(string, interface{})

//按字符顺序
func traverseMapInStringOrder(params map[string]interface{}, handler mapEntryHandler) {
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		handler(k, params[k])
	}
}

//MD5
func generateMd5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// val转str
func valToStr(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

// 签名
func sign(obj interface{}, token string) (string, error) {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		return "", err
	}
	origin := ""
	traverseMapInStringOrder(data, func(key string, value interface{}) {
		if strings.ToLower(key) != "sign" {
			origin += valToStr(value)
		}
	})
	origin += token
	return generateMd5(origin), nil
}

// 签名body数据
func signBody(body interface{}, token string) (map[string]interface{}, error) {
	info := map[string]interface{}{}
	if body == nil {
		return nil, errors.New("empty body")
	}
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonStr, &info)
	if err != nil {
		return nil, err
	}
	info["timestamp"] = time.Now().Unix()
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	info["nonce"] = strings.Replace(uid.String(), "-", "", -1)
	s, err := sign(info, token)
	if err != nil {
		return nil, err
	}
	info["sign"] = s
	return info, nil
}
