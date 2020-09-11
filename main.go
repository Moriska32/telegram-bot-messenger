package main

import (
	telebot "telegram-bot-messenger/Telebot"
)

func main() {
	bot := telebot.BotINIT()

	telebot.SendMessegeBot(bot, "privat", "Hallo")

}
