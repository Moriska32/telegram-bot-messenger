package telebot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"telegram-bot-messenger/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type configJSON struct {
	BotKey string `json:"bot-key"`
}

//TelebotInit initialization bot
func TelebotInit() *tgbotapi.BotAPI {

	dbConnect := config.Connect()
	defer dbConnect.Close()

	jsonFile, err := os.Open("config.json")
	defer jsonFile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	var config configJSON

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)

	bot, err := tgbotapi.NewBotAPI(config.BotKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		query := "select from nami.fn_telegram_ins(?,?)"

		_, err := dbConnect.Exec(query, update.ChannelPost.Chat.ID, update.ChannelPost.Chat.Type)

		if err != nil {
			log.Panic(err.Error())
		}
	}
	return bot
}

//SendMessegeBot Call for sand messege to somebody
func SendMessegeBot(t *tgbotapi.BotAPI) {

	dbConnect := config.Connect()
	defer dbConnect.Close()

	query := "select from nami.fn_telegram_sel(?)"
	_ = query

}
