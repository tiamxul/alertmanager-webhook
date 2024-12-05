package model

type DDMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown Md     `json:"markdown,omitempty"`
}

type Md struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	IsAtAll   bool     `json:"isAtAll"`
	AtMobiles []string `json:"atMobiles,omitempty"`
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
