# 简介

针对接入 https://app.qubic.li/ 进行挖矿的节点，通过 API 监控 Qubic 节点运行情况，并发送微信告警

通过提供钱包地址和本钱包地址下运行的节点数量，来跟官方 API 获取到信息进行对比，如果 **实际运行节点数量 > 活跃节点数量** 则发送微信告警

程序每60秒执行一次检测



# 使用方法

## 1.在程序的同目录中准备配置文件  `qubic.json` 

参数说明：

- `weixin` 在下边 **附录：获取微信机器人 Webhook** 中有获取链接说明
- Name: 钱包名称
- Count: 本钱包一共有多少个节点在运行
- Wallet: 钱包地址
- **修改配置文件不用重启程序，会自动检测最新配置**

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

```cmd
@echo off

:LOOP
REM 检查进程是否存在
tasklist | find /i "monitor_qubic-windows-amd64.exe"

REM 如果不存在，则启动你的脚本或命令
if errorlevel 1 (
     monitor_qubic-windows-amd64.exe
)

REM 输出当前系统时间
echo Current system time: %time%sss

REM 延时5秒
timeout /t 10 >nul

REM 跳转回循环开始
goto LOOP
```



## 3.Linux 下以 Systemd 方式运行

```sh
## 1.下载编译好的可执行文件或者自己编译，并将问文件放入到 /usr/local/bin/ 目录下并给予可执行权限
## https://github.com/xiaoliuxiao6/monitor_qubic/releases

## 准备服务文件
cat <<\EOF >/etc/systemd/system/monitor_qubic.service
[Unit]
Description=监控指定目录下最后一个文件更新时间是否超过指定时间并发送微信告警

[Service]
Restart=always
RestartSec=5
ExecStart=/usr/local/bin/monitor_qubic-linux-amd64

[Install]
WantedBy=multi-user.target
EOF

## 启动并设置为开机自动启动
systemctl enable monitor_qubic.service
systemctl stop monitor_qubic.service
systemctl start monitor_qubic.service
systemctl status monitor_qubic.service
```





# 附录：获取微信机器人 Webhook

## 1.登录手机企业微信

## 2.打开群聊

- 右上角的三个点 - 群机器人 - 添加（取个名字）
- 复制 Webhook 地址

