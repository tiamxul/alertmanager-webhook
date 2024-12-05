package main

import (
	"alertmanager-webhook/log"

	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ENV           string `yaml:"env"`
	LogLevel      string `yaml:"log_level"`
	ListenAddress string `yaml:"listen_address"`
	HotReload     bool   `yaml:"hot_reload"`
	AlertType     string `yaml:"alert_type"`
	OpenDingding  string `yaml:"open_dingding"`
	OpenFeishu    string `yaml:"open_feishu"`
	Dingding      `yaml:"dingding"`
	Feishu        `yaml:"feishu"`
}
type Dingding struct {
	DdUrl string `yaml:"dd_url"`
}
type Feishu struct {
	FsUrl string `yaml:"fs_url"`
}

const configPath = "config/config.yaml"

func (c *Config) Initial() (err error) {
	defer func() {
		if err == nil {
			log.Printf("config initialed, env: %s", cfg.ENV)
		}
	}()

	if level, err := logrus.ParseLevel(c.LogLevel); err != nil {
		return err
	} else {
		log.DefaultLogger().SetLevel(level)
	}
	return nil
}
func loadConfig() {
	cfg = new(Config)
	multiconfig.MustLoadWithPath(configPath, cfg)
}
