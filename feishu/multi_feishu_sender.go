package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"alertmanager-webhook/model"
)

type MultiFeiShuSender struct {
	Robots []RobotConfig
}

type RobotConfig struct {
	Name       string `json:"name"`
	WebhookURL string `json:"webhook_url"`
}

func (m *MultiFeiShuSender) SendToRobot(name string, message *model.CommonMessage) error {
	for _, robot := range m.Robots {
		if robot.Name == name {
			var msg interface{}
			if message.Title == "" {
				msg = model.NewTextMessage(message.Text)
			} else {
				content := [][]model.PostMessageContentPostZhCnContent{
					{
						*model.NewPostMessageContentPostZhCnContent("text", message.Text, "", "", "", "", "", ""),
					},
				}
				msg = model.NewPostMessage(message.Title, content)
			}

			payload, err := json.Marshal(msg)
			if err != nil {
				return err
			}

			resp, err := http.Post(robot.WebhookURL, "application/json", bytes.NewBuffer(payload))
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			fmt.Printf("FeiShu response from %s: %s\n", name, body)
			return nil
		}
	}
	return fmt.Errorf("robot not found: %s", name)
}
