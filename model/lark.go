package model

// 定义飞书消息类型
type FSMessage struct {
	Msgtype string  `json:"msg_type"`
	Content Content `json:"content"`
}
type Content struct {
	Text string `json:"text"`
}

// 响应体相关
type LarkResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
type Data struct {
}

// {
//     "msg_type": "text",
//     "content": {
//         "text": "新更新提醒"
//     }
// }
