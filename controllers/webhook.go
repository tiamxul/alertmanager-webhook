package controllers

import (
	"net/http"
	"path/filepath"

	"alertmanager-webhook/config"
	"alertmanager-webhook/interfaces"
	"alertmanager-webhook/log"
	"alertmanager-webhook/model"
	"alertmanager-webhook/senders/dingtalk"
	"alertmanager-webhook/senders/feishu"

	"github.com/gin-gonic/gin"
)

func HandlerWebhook(c *gin.Context) {
	var notification model.Alert
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	templateName := getTemplateName(notification.Labels) // Implement logic to choose a template based on labels
	templateFile := filepath.Join("templates", templateName+".tmpl")
	alertTemplate, err := model.NewTemplate(templateFile)
	if err != nil {
		log.Errorf("Failed to load template file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "template loading failed"})
		return
	}

	// Apply the template to generate the message content
	messageContent, err := alertTemplate.Execute(notification)
	if err != nil {
		log.Errorf("Failed to execute template: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "template execution failed"})
		return
	}

	notification.Text = messageContent

	sender := getSender(notification)
	if sender == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported platform"})
		return
	}

	err = sender.Send(&notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful send alert notification!"})
}

func getSender(notification model.CommonMessage) interfaces.MessageSender {
	switch notification.Platform {
	case "dingtalk":
		if config.AppConfig.Dingtalk.Open == 1 {
			return &dingtalk.DingTalkSender{WebhookURL: config.AppConfig.Dingtalk.WebhookURL}
		}
	case "feishu":
		if config.AppConfig.OpenFeishu == 1 {
			robotName := getFeishuRobotName(notification.Labels) // Implement logic to choose a robot based on labels
			return &feishu.MultiFeiShuSender{Robots: config.AppConfig.FeishuRobots}
		}
	}

	return nil
}

func getTemplateName(labels map[string]string) string {
	// Implement logic to choose a template based on labels
	// This is an example implementation and should be customized according to your needs.
	if val, ok := labels["template"]; ok {
		return val
	}
	return "default" // Default template name
}

func getFeishuRobotName(labels map[string]string) string {
	// Implement logic to choose a robot based on labels
	// This is an example implementation and should be customized according to your needs.
	if val, ok := labels["robot"]; ok {
		return val
	}
	return "robot1" // Default robot name
}
