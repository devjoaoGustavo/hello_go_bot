package main

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
