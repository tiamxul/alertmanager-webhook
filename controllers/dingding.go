package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tiamxu/alertmanager-webhook/log"
	"github.com/tiamxu/alertmanager-webhook/model"
)

func PostToDingding(title, text, Ddurl, open string) string {
	if open != "1" {
		log.Infoln("钉钉接口未配置未开启状态,请先配置open-dingding为1")
		return "钉钉接口未配置未开启状态,请先配置open-dingding为1"
	}
	Atall := false
	atMobile := []string{"15888888888"}
	SendText := text
	u := model.DDMessage{
		Msgtype: "markdown",
		Markdown: model.Md{
			Title: title,
			Text:  SendText,
		},
		At: model.At{
			AtMobiles: atMobile,
			IsAtAll:   Atall,
		},
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(u)
	if err == nil {
		fmt.Print("json.NewEncoder 编码结果: ", b.String())
	}
	log.Infoln("[dingding]", b)
	// client := &http.Client{}
	resp, err := http.Post(Ddurl, "application/json", b)
	if err != nil {
		log.Errorln("[dingding]", err.Error())
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorln("[dingding]", err.Error())
	}
	log.Infoln("[dingding]", string(result))
	return string(result)

}
func TransformToMarkdown(notification model.AlertMessage) (markdown *model.DDMessage, err error) {

	groupKey := notification.GroupKey
	status := notification.Status

	// annotations := notification.CommonAnnotations

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))
	buffer.WriteString(fmt.Sprintf("#### 告警项: \n"))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	}

	markdown = &model.DDMessage{
		Msgtype: "markdown",
		Markdown: model.Md{
			Title: fmt.Sprintf("通知组：%s(当前状态:%s)", groupKey, status),
			Text:  buffer.String(),
		},
		At: model.At{
			IsAtAll: false,
		},
	}

	return
}
