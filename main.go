package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const telegramUrl = "https://api.telegram.org/"
const token = "<token>"

const getMePath = "getme"
const getUpdatesPath = "getupdates"
const sendMessagePath = "sendmessage"

type Bot struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName  string `json:"last_name`
}

type ResponseBot struct {
	Ok     bool `json:"ok"`
	Result Bot  `json:"result"`
}

func main() {
	var body []byte
	var bot ResponseBot

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, getMePath)

	resp, err := http.Get(url)
	if err != nil {
		s := err.Error()
		fmt.Println(s)
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &bot)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bot)
	fmt.Println("ok:", bot.Ok)
	fmt.Println("result:", bot.Result)
	fmt.Println("Id:", bot.Result.Id)
	fmt.Println("IsBot:", bot.Result.IsBot)
	fmt.Println("FirstName:", bot.Result.FirstName)
	fmt.Println("UserName:", bot.Result.UserName)
}
