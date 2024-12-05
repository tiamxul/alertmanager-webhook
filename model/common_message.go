package model

// CommonMessage is a generic message structure that can be sent through various channels.
type CommonMessage struct {
	Text     string `json:"text"`
	Title    string `json:"title,omitempty"`
	Platform string `json:"platform"` // "dingtalk" or "feishu"
}
