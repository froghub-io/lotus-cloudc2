
一个基于filecoin lotus优化的一个云c2的实现，仅需要改动一小部分代码，就能实现c2远端执行。 

Cloud C2 有三部分组成：
- 本项目是Cloud C2的客户端实现
- [Froghub控制台](https://console.froghub.cn/) 解决用户体验问题，包括存储服务管理，设备管理和费用管理
- [cloudc2-deamon](https://github.com/froghub-io/cloudc2-daemon) 优化过的c2执行worker，有守护进程解决更新和动态支持32G和62G扇区

适用人群：
- 正在自己优化lotus，c2未优化或者比本方案慢。
- 有空闲的封装机，希望创造收入
- 正在封装的集群，没有c2机器，或者c2计算能力不足。

## 合并lotus-cloudc2到自己的优化版本

默认master分支是dev分支，是编写说明文档的分之，建议使用发行版本对应的分支。

仅需要合并名为"_**实现clondc2,需要合并的代码**_"的提交

目前支持版本如下
- [v1.13.0](https://github.com/froghub-io/lotus-cloudc2/tree/cloudc2/v1.13.0)
- [v1.12.0](https://github.com/froghub-io/lotus-cloudc2/tree/cloudc2/v1.12.0)
- [~~v1.11.2~~](https://github.com/froghub-io/lotus-cloudc2/tree/cloudc2/v1.11.2)
- [~~1.11.1~~](https://github.com/froghub-io/lotus-cloudc2/tree/cloudc2/v1.11.1)
- [~~v1.11.0~~](https://github.com/froghub-io/lotus-cloudc2/tree/cloudc2/v1.11.0)

## 接入使用
正式环境请进入控制台申请token，并设置环境变量
```shell
export CLOUD_C2_TOKEN=您的token
```

## 报告漏洞

请发送电子邮件至froghub@163.com。

## 联系我们
- [微信](https://twitter.com/filecoin)

<img src="https://storageapi.fleek.co/froghubman-team-bucket/5df07fba-972f-4c66-9089-00c4bd007768.png" width="200" />

- Email：  froghub@163.com

## License

Dual-licensed under [MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
