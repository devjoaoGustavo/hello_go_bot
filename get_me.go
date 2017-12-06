package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const getMePath = "getMe"

type GetMe struct {
	Ok     bool
	Result User
}

func showBotInfo(config Config) {
	var body []byte
	var getMe GetMe
	url := fmt.Sprintf(
		"%s/bot%s/%s",
		config.HelloGoBot.Telegram.Api.EndPoint,
		config.HelloGoBot.Telegram.Api.Token,
		getMePath,
	)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &getMe)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Esta é a estrutura da resposta do método getMe")
	fmt.Println("ok:", getMe.Ok)
	fmt.Println("result:", getMe.Result)
	fmt.Println("User Id:", getMe.Result.Id)
	fmt.Println("User IsBot:", getMe.Result.IsBot)
	fmt.Println("User FirstName:", getMe.Result.FirstName)
	fmt.Println("User LastName:", getMe.Result.LastName)
	fmt.Println("User UserName:", getMe.Result.UserName)
}
