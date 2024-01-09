package model

// 定义钉钉消息类型

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}
type Md struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DDMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown Md     `json:"markdown"`
	At       At     `json:"at"`
}

// {
// 	"msgtype": "markdown",
// 	"markdown": {
// 		"title":"Prometheus告警信息",
// 		"text": "#### 监控指标\n" +
// 				"> 监控描述信息\n\n" +
// 				"> ###### 告警时间 \n"
// 	},
//    "at": {
// 	   "atMobiles": [
// 		   "156xxxx8827",
// 		   "189xxxx8325"
// 	   ],
// 	   "isAtAll": false
//    }
// }
