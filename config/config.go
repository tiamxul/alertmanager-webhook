package config

import (
	"fmt"
	"log"
	"os"

	"github.com/tiamxu/alertmanager-webhook/model"
	"gopkg.in/yaml.v3"
)

var Config *model.Config

func init() {
	fmt.Println("####init")
	filename := "./config/config.yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("yaml文件内容:\n%v\n", string(data))
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", Config)
}
