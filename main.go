package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	HelloGoBot struct {
		Telegram struct {
			Api struct {
				EndPoint string
				Token    string
			}
		}
	} `yaml:"hello_go_bot"`
}

func main() {
	var config Config
	configFile, err0 := os.Open("./config.yml")
	if err0 != nil {
		s := err0.Error()
		fmt.Println(s)
		return
	}
	configBytes, err1 := ioutil.ReadAll(configFile)
	if err1 != nil {
		s := err1.Error()
		fmt.Println(s)
		return
	}
	err2 := yaml.Unmarshal(configBytes, &config)
	if err2 != nil {
		s := err2.Error()
		fmt.Println(s)
		return
	}
	showBotInfo(config)
	fmt.Println()
	showUpdates(config)
}
