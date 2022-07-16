package main

import (
	"github.com/andrew528i/hardhat-gateway/api"
	"github.com/andrew528i/hardhat-gateway/bot"
)

func main() {
	done := make(chan bool)
	messages := make(chan string)

	go api.Run(done, messages)
	go bot.Run(done, messages)

	<-done
}
