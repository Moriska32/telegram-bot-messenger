package main

import (
	"log"

	telebot "github.com/Moriska32/telegram-bot-messenger/Telebot"
)

func main() {
	//Создание бота
	bot := telebot.BotINIT()
	var (
		who  string = "channel"
		text string = "hallo"
	)
	//Отправка сообщений channel - В чаты и каналы, private - в личные сообщения
	err := telebot.SendMessegeBot(bot, who, text)
	if err != nil {
		log.Println(err)
	}

	var forever chan string
	<-forever

}
