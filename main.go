package main

import (
	"fmt"

	"github.com/Tony-Moon/CasinoCord/bot"
	"github.com/Tony-Moon/CasinoCord/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println("Failed to configure bot.")
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
