package controllers

//curl -X POST -H "Content-Type: application/json" \
//    -d '{"msg_type":"text","content":{"text":"request example"}}' \
//   https://open.feishu.cn/open-apis/bot/v2/hook/****

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tiamxu/alertmanager-webhook/log"
	"github.com/tiamxu/alertmanager-webhook/model"
)

func TransformToLarkRequest(notification model.AlertMessage) (larkRequest model.FSMessage, err error) {
	var buffer bytes.Buffer

	// 先拿到分组情况
	buffer.WriteString(fmt.Sprintf("通知组%s,状态[%s]\n告警项\n\n", notification.GroupKey, notification.Status))

	// 每条告警逐个获取，拼接到一起
	for _, alert := range notification.Alerts {
		buffer.WriteString(fmt.Sprintf("摘要：%s\n详情: %s\n", alert.Annotations["summary"], alert.Annotations["description"]))
		buffer.WriteString(fmt.Sprintf("开始时间: %s\n\n", alert.StartsAt.Format("15:04:05")))
	}

	// 构造出飞书机器人所需的数据结构
	larkRequest = model.FSMessage{
		Msgtype: "text",
		Content: model.Content{
			Text: buffer.String(),
		},
	}

	return larkRequest, nil
}
func PostToFS(text, Fsurl, open string) string {
	if open != "1" {
		log.Infoln("[feishu]", "飞书接口未配置未开启状态,请先配置open-feishu为1")
		return "飞书接口未配置未开启状态,请先配置open-feishu为1"
	}
	RTstring := ""
	if strings.Contains(Fsurl, "/v2/") {
		RTstring = PostToFeiShu(text, Fsurl)
	} else {
		RTstring = PostToFeiShu(text, Fsurl)
	}
	return RTstring
}
func PostToFeiShu(text, Fsurl string) string {
	u := model.FSMessage{
		Msgtype: "text",
		Content: model.Content{
			Text: text,
		},
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(u)
	if err == nil {
		fmt.Print("json.NewEncoder 编码结果: ", b.String())
	}
	log.Infoln("[feishu]", b)
	// client := &http.Client{}
	resp, err := http.Post(Fsurl, "application/json", b)
	if err != nil {
		log.Errorln("[feishu]", err.Error())
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorln("[feishu]", err.Error())
	}
	log.Infoln("[feishu]", string(result))

	return string(result)

}
