# 简介

通过 API 监控 Qubic 节点运行情况，并发送微信告警

通过提供钱包地址和本钱包地址下运行的节点数量，来跟官方 API 获取到信息进行对比，如果 **实际运行节点数量 > 活跃节点数量** 则发送微信告警





# 使用方法

## 1.在当前目录准备配置文件  `qubic.json`

参数说明：

- `weixin` 在下边 **附录：获取微信机器人 Webhook** 中有获取链接说明
- Name: 钱包名称
- Count: 本钱包一共有多少个节点在运行
- Wallet: 钱包地址

```json
{
  "weixin": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=d55xxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
  "Nodes":[
    {
      "Name":"qianbao1",
      "Count":1,
      "wallet":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    },
    {
      "Name":"qianbao2",
      "Count":18,
      "Wallet":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    },
    {
      "Name":"qianbao3",
      "Count":48,
      "wallet":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    }
  ]
}
```



## 2.Windows 运行

将以下信息保存为 `.cmd` 或 `.bat` 文件双击运行即可

```

```



## 3.Linux 运行

```

```





# 附录：获取微信机器人 Webhook

## 1.登录手机企业微信

## 2.打开群聊

- 右上角的三个点 - 群机器人 - 添加（取个名字）
- 复制 Webhook 地址

