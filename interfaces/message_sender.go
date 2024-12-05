package interfaces

import "alertmanager-webhook/model"

type MessageSender interface {
	Send(message *model.CommonMessage) error
}
