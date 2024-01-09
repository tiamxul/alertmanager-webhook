package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiamxu/alertmanager-webhook/config"
	"github.com/tiamxu/alertmanager-webhook/log"
	"github.com/tiamxu/alertmanager-webhook/model"
)

func SendMessage(message model.AlertMessage, ddurl string) (err error) {
	markdown, err := TransformToMarkdown(message)
	if err != nil {
		return
	}
	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", ddurl, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return
	}
	fmt.Printf("response body: %v\n", string(body))
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	return
}

var url = "https://oapi.dingtalk.com/robot/send?access_token=xxxxx"

func HandlerWebhook(c *gin.Context) {
	var notification model.AlertMessage
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	err = SendMessage(notification, url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})
}

const (
	LARK_URL = "https://open.feishu.cn/open-apis/bot/v2/hook/bf8bb912-bc2e-40ad-9533-fcb8068aa621"
)

func FeishuAlertmanagerWebhook(c *gin.Context) {
	var notification model.AlertMessage

	// 绑定对象
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Infof("收到alertmanager告警:\n%s", notification)

	// 根据alertmanager的请求构造飞书消息的请求数据结构
	larkRequest, _ := TransformToLarkRequest(notification)

	// 向飞书服务器发送POST请求，将飞书服务器返回的内容转为对象
	bytesData, _ := json.Marshal(larkRequest)
	req, _ := http.NewRequest("POST", LARK_URL, bytes.NewReader(bytesData))
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	// 飞书服务器可能通信失败
	if err != nil {
		log.Errorf("请求飞书服务器失败：%s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var larkResponse model.LarkResponse
	err = json.Unmarshal([]byte(body), &larkResponse)
	// 飞书服务器返回的包可能有问题
	if err != nil {
		log.Errorf("获取飞书服务器响应失败：%s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}
	log.Infoln("向飞书服务器发送消息成功")
	c.JSON(http.StatusOK, gin.H{
		"message": "successful receive alert notification message!",
	})
}
func PrometheusAlert(c *gin.Context) {
	var notification model.AlertMessage
	// 绑定对象
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Infof("收到alertmanager告警:\n%s", notification)
	var fstext string
	fstext = "###[通知组]" + notification.GroupKey + "[状态]" + notification.Status + "\n" + "[告警项]" + "\n\n"

	for _, alert := range notification.Alerts {
		fstext = fstext + "####[摘要]:" + alert.Annotations["summary"] + "\n" + "####[详情:]" + alert.Annotations["description"] + "\n"
	}
	// var open string = "1"
	open := config.Config.OpenFeishu
	fmt.Println("###open:", open)
	PostToFS(fstext, LARK_URL, open)
	//PostToFeiShu(fstext, LARK_URL)
}
