/*
* @Descripttion:
* @Author: 1327133225@qq.com
* @version:
* @Date: 2023-11-22 10:34:29
* @LastEditors: 1327133225@qq.com
* @LastEditTime: 2023-11-22 10:48:09
 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// 解析钱包信息
	var nodes []Node
	if err := viper.UnmarshalKey("Nodes", &nodes); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	// 通过钱包地址获取算力信息
	for _, node := range nodes {
		activeNodeCount, totalIts := getPerformance(node)
		info := fmt.Sprintf("节点名称: %v, 节点总数: %v, 活跃节点数: %v, 总算力: %v",
			node.Name, node.Count, activeNodeCount, totalIts)

		// 如果节点数量 > 活跃节点数即是有节点掉线，发送微信告警
		if node.Count > activeNodeCount {
			SendWeixin("Qubic 有节点掉线：" + info)
			log.Warnln(info)
		} else {
			log.Println(info)
		}

		// 早8点和晚上8点发送存活告警
		hostname, _ := os.Hostname()
		if time.Now().Minute() < 10 {
			if time.Now().Hour() == 8 || time.Now().Hour() == 20 {
				SendWeixin(fmt.Sprintf("我正在监控 Qubic 的运行情况，正常运行中。。。", hostname))
			}
		}

	}
}

func SendWeixin(message string) {
	// WeixinWebhook := viper.GetString("Channel.Weixin.Webhook")
	// 1. https://github.com/feiyu563/PrometheusAlert/blob/master/doc/readme/conf-wechat.md
	// 2. 登录企业微信网页 - 我的企业 - 微信插件 - 扫描二维码即可手机微信来接收信息

	url := viper.GetString("weixin")

	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{
    	"msgtype": "text",
    	"text": {
        	"content": "%v"
    	}
	}`, message))
	client := &http.Client{}
	// req, err := http.NewRequest(method, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=0da619f2-afbf-4d2f-a13c-db23abba986e", payload)
	//req, err := http.NewRequest(method, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ff8d99d2-5d81-4d20-815f-b55d3ffa0d59", payload)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatalf("发送微信信息失败1: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("发送微信信息失败2: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// 通过 API 来获取当前节点算力情况
func getPerformance(node Node) (activeNodeCount int, totalIts int) {

	activeNodeCount = 0 // 算力 > 0 的节点
	totalIts = 0        // 总算力

	URL := fmt.Sprintf("https://api.qubic.li/PublicPool/Performance/%v", node.Wallet)

	// 1.定义一个 Request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
	}

	// 2.为此 Request 调用 http.DefaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}

	// 3.延迟关闭
	defer resp.Body.Close()

	// 读取结果
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取结果失败: %v", err)
		return
	}

	// 转换为结构体
	var performances Performances
	if err := json.Unmarshal(body, &performances); err != nil {
		log.Fatalf("解析到结构体失败: %v", err)
	}

	for _, performance := range performances {
		if performance.CurrentIts > 0 {
			activeNodeCount += 1
			totalIts += performance.CurrentIts
		}
	}

	return activeNodeCount, totalIts
}

type Performances []struct {
	ID              string  `json:"id"`
	MinerBinaryID   any     `json:"minerBinaryId"`
	Alias           string  `json:"alias"`
	LastActive      string  `json:"lastActive"`
	CurrentIts      int     `json:"currentIts"`
	CurrentIdentity any     `json:"currentIdentity"`
	SolutionsFound  int     `json:"solutionsFound"`
	Threads         any     `json:"threads"`
	TotalFeeTime    float64 `json:"totalFeeTime"`
	FeeReports      []any   `json:"feeReports"`
	IsActive        bool    `json:"isActive"`
}

// 从配置文件读取节点信息
func init() {
	viper.SetConfigName("qubic")          // 配置文件的文件名(不带扩展名)
	viper.SetConfigType("json")           // (文件的扩展名)如果 SetConfigName 中没有扩展名，则需要
	viper.AddConfigPath("./")             // 配置文件路径(可以多次出现以搜索多个路径)
	viper.AddConfigPath("/usr/local/etc") // 配置文件路径(可以多次出现以搜索多个路径)

	err := viper.ReadInConfig() // 载入配置文件
	if err != nil {
		log.Panicf("配置文件读取失败: %v \n", err)
	} else {
		log.Printf("配置文件读取成功: %v", viper.ConfigFileUsed())
	}

	viper.WatchConfig() //监视和重新读取配置文件
}

// Node 结构体表示配置文件中的一个节点
type Node struct {
	Name   string
	Count  int
	Wallet string
}