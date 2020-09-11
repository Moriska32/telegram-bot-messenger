package main

import (
	"log"

	telebot "github.com/Moriska32/telegram-bot-messenger/Telebot"
)

func main() {

	bot := telebot.BotINIT()
	var (
		who  string = "channel"
		text string = "hallo"
	)
	err := telebot.SendMessegeBot(bot, who, text)
	if err != nil {
		log.Println(err)
	}

	var forever chan string
	<-forever

}
