package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const getUpdatesPath = "getUpdates"

/*
TODO: Adicionar todas as propriedades descritas na documentação da API
https://core.telegram.org/bots/api#getupdates
*/

type Entity struct {
	OffSet int `json:"off_set"`
	Length int
	Type   string
}
type Chat struct {
	Id        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name`
	UserName  string `json:"user_name"`
	Type      string
}
type From struct {
	Id           int
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name`
	UserName     string `json:"user_name"`
	LanguageCode string `json:"language_code"`
}
type Message struct {
	MessageId int `json:"message_id"`
	From      From
	Chat      Chat
	Date      int
	Text      string
	Entities  []Entity
}
type Update struct {
	UpdateId int `json:"update_id"`
	Message  Message
}
type GetUpdates struct {
	Ok     bool
	Result []Update
}

func showUpdates(config Config) {
	var updates GetUpdates
	url := fmt.Sprintf("%s/bot%s/%s", config.HelloGoBot.Telegram.Api.EndPoint, config.HelloGoBot.Telegram.Api.Token, getUpdatesPath)
	resp, err0 := http.Get(url)
	if err0 != nil {
		s := err0.Error()
		fmt.Printf("%q\n", s)
		return
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		s := err1.Error()
		fmt.Printf("%q\n", s)
		return
	}
	err2 := json.Unmarshal(body, &updates)
	if err2 != nil {
		s := err2.Error()
		fmt.Printf("%q\n", s)
		return
	}
	fmt.Println("Esta é a resposta do método getUpdates")
	fmt.Println("Ok:", updates.Ok)
	fmt.Println("result:", updates.Result)

	for _, update := range updates.Result {
		fmt.Println("UpdateID:", update.UpdateId)
		fmt.Println("Message:", update.Message)
		fmt.Println("Message MessageId:", update.Message.MessageId)
		fmt.Println("Message From:", update.Message.From)
		fmt.Println("Message From Id:", update.Message.From.Id)
		fmt.Println("Message From IsBot:", update.Message.From.IsBot)
		fmt.Println("Message From FirstName:", update.Message.From.FirstName)
		fmt.Println("Message From LastName:", update.Message.From.LastName)
		fmt.Println("Message From UserName:", update.Message.From.UserName)
		fmt.Println("Message Chat:", update.Message.Chat)
		fmt.Println("Message Chat Id:", update.Message.Chat.Id)
		fmt.Println("Message Chat FirstName:", update.Message.Chat.FirstName)
		fmt.Println("Message Chat LastName:", update.Message.Chat.LastName)
		fmt.Println("Message Chat UserName:", update.Message.Chat.UserName)
		fmt.Println("Message Chat Type:", update.Message.Chat.Type)
		fmt.Println("Message Date:", update.Message.Date)
		fmt.Println("Message Text:", update.Message.Text)
		fmt.Println("Message Entities:", update.Message.Entities)
		for _, entity := range update.Message.Entities {
			fmt.Println("Message Entities OffSet:", entity.OffSet)
			fmt.Println("Message Entities Length:", entity.Length)
			fmt.Println("Message Entities Type:", entity.Type)
		}
	}
}
