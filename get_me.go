package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const getMePath = "getMe"

/*
TODO: Adicionar todas as propriedades descritas na documentação da API
https://core.telegram.org/bots/api#getme
*/

type GetMe struct {
	Ok     bool `json:"ok"`
	Result struct {
		Id        int64  `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
		UserName  string `json:"last_name`
	}
}

func showBotInfo(config Config) {
	var body []byte
	var getMe GetMe
	url := fmt.Sprintf("%s/bot%s/%s", config.HelloGoBot.Telegram.Api.EndPoint, config.HelloGoBot.Telegram.Api.Token, getMePath)
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

	err = json.Unmarshal(body, &getMe)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Esta é a estrutura da resposta do método getMe")
	fmt.Println("ok:", getMe.Ok)
	fmt.Println("result:", getMe.Result)
	fmt.Println("Id:", getMe.Result.Id)
	fmt.Println("IsBot:", getMe.Result.IsBot)
	fmt.Println("FirstName:", getMe.Result.FirstName)
	fmt.Println("UserName:", getMe.Result.UserName)
}
