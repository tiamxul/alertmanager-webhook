package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env           string           `yaml:"env"`
	LogLevel      string           `yaml:"log_level"`
	ListenAddress string           `yaml:"listen_address"`
	HotReload     bool             `yaml:"hot_reload"`
	AlertType     string           `yaml:"alert_type"`
	Dingtalk      DingtalkConfig   `yaml:"dingtalk"`
	OpenFeishu    int              `yaml:"open_feishu"`
	FeishuRobots  []RobotConfig    `yaml:"feishu_robots"`
	Templates     []TemplateConfig `yaml:"templates"`
}

type DingtalkConfig struct {
	Open       int    `yaml:"open_dingtalk"`
	WebhookURL string `yaml:"webhook_url"`
}

type RobotConfig struct {
	Name       string `yaml:"name"`
	WebhookURL string `yaml:"webhook_url"`
}

type TemplateConfig struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

var AppConfig *Config

func Load() error {
	filename := "./config/default.yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	AppConfig = &Config{}
	err = yaml.Unmarshal(data, AppConfig)
	if err != nil {
		return err
	}

	return nil
}
