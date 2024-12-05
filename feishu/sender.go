package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"alertmanager-webhook/model"
)

type FeiShuSender struct {
	WebhookURL string
}

func (f *FeiShuSender) Send(message *model.CommonMessage) error {
	switch message.Platform {
	case "feishu":
		var msg interface{}
		if message.Title == "" {
			msg = model.NewTextMessage(message.Text)
		} else {
			content := [][]model.PostMessageContentPostZhCnContent{
				{
					*model.NewPostMessageContentPostZhCnContent("text", message.Text, "", "", "", "", "", "")},
			}
			msg = model.NewPostMessage(message.Title, content)
		}

		payload, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		resp, err := http.Post(f.WebhookURL, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Printf("FeiShu response: %s\n", body)
		return nil
	default:
		return fmt.Errorf("unsupported platform for FeiShuSender")
	}
}
