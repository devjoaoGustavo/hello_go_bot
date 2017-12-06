package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	var config Config
	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	showBotInfo(config)
	fmt.Println()
	showUpdates(config)
}
