package model

type MsgType string

const (
	MsgTypeText MsgType = "text"
	MsgTypePost MsgType = "post"
	// MsgTypeImage       MsgType = "image"
	// MsgTypeShareChat   MsgType = "share_chat"
	// MsgTypeInteractive MsgType = "interactive"
)

// 简单文本消息
type TextMessage struct {
	MsgType MsgType            `json:"msg_type"`
	Content TextMessageContent `json:"content"`
}

type TextMessageContent struct {
	Text string `json:"text"`
}

func NewTextMessage(text string) *TextMessage {
	return &TextMessage{
		MsgType: MsgTypeText,
		Content: TextMessageContent{
			Text: text,
		},
	}
}

// 富文本消息
type PostMessage struct {
	MsgType MsgType            `json:"msg_type"`
	Content PostMessageContent `json:"content"`
}

type PostMessageContent struct {
	Post PostMessageContentPost `json:"post"`
}

type PostMessageContentPost struct {
	ZhCn PostMessageContentPostZhCn `json:"zh-CN"`
}

type PostMessageContentPostZhCn struct {
	Title   string                                `json:"title"`
	Content [][]PostMessageContentPostZhCnContent `json:"content"`
}

type PostMessageContentPostZhCnContent struct {
	Tag       string `json:"tag"`
	Text      string `json:"text,omitempty"`
	Href      string `json:"href,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	ImageKey  string `json:"image_key,omitempty"`
	FileKey   string `json:"file_key,omitempty"`
	EmojiType string `json:"emoji_type,omitempty"`
}

func NewPostMessageContentPostZhCnContent(tag, text, href, userId, userName, imageKey, fileKey, emojiType string) *PostMessageContentPostZhCnContent {
	return &PostMessageContentPostZhCnContent{
		Tag:       tag,
		Text:      text,
		Href:      href,
		UserId:    userId,
		UserName:  userName,
		ImageKey:  imageKey,
		FileKey:   fileKey,
		EmojiType: emojiType,
	}
}

func NewPostMessage(title string, content [][]PostMessageContentPostZhCnContent) *PostMessage {
	return &PostMessage{
		MsgType: MsgTypePost,
		Content: PostMessageContent{
			Post: PostMessageContentPost{
				ZhCn: PostMessageContentPostZhCn{
					Title:   title,
					Content: content,
				},
			},
		},
	}
}
