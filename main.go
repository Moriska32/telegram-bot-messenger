package main

import (
	telebot "telegram-bot-messenger/Telebot"
)

func main() {
	bot := telebot.TelebotInit()
	telebot.SendMessegeBot(bot)

}
