package telebot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Moriska32/telegram-bot-messenger/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type configJSON struct {
	BotKey string `json:"bot-key"`
}

//BotINIT initialization bot
func BotINIT() *tgbotapi.BotAPI {

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

	go func(tgbotapi.UpdatesChannel) {

		for update := range updates {

			if update.Message != nil {

				err := loadtodb(update.Message.Chat.ID, "NULL", update.Message.Chat.Type)

				if err != nil {
					log.Println(err)
				}

			}

			if update.ChannelPost != nil {
				err := loadtodb(update.ChannelPost.Chat.ID, update.ChannelPost.Chat.Title, update.ChannelPost.Chat.Type)

				if err != nil {
					log.Println(err)
				}

			}
		}
	}(updates)
	return bot
}

func loadtodb(id int64, title string, who string) error {

	dbConnect := config.Connect()
	defer dbConnect.Close()

	query := fmt.Sprintf("select from nami.fn_telegram_ins(%d,'%s','%s')", id, title, who)

	_, err := dbConnect.Exec(query)

	return err
}

//UserJSON json
type UserJSON []struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

//SendMessegeBot Call for sand messege to somebody
func SendMessegeBot(t *tgbotapi.BotAPI, who string, text string) error {

	dbConnect := telebot.config.Connect()
	defer dbConnect.Close()

	var (
		pool string
	)
	todo := fmt.Sprintf("SELECT nami.fn_telegram_sel('%s');", who)
	_ = todo
	sql, err := dbConnect.Query(todo)

	if err != nil {
		return err
	}
	for sql.Next() {
		var user UserJSON

		sql.Scan(&pool)
		log.Printf(pool)
		err = json.Unmarshal([]byte(pool), &user)

		if err != nil {
			return err
		}
		for _, i := range user {
			msg := tgbotapi.NewMessage(int64(i.ID), text)
			t.Send(msg)
		}
	}
	return err
}
