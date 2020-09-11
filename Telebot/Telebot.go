package telebot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"telegram-bot-messenger/config"

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
	return bot
}

func loadtodb(id int64, title string, who string) error {

	dbConnect := config.Connect()
	defer dbConnect.Close()

	query := fmt.Sprintf("select from nami.fn_telegram_ins(%d,'%s','%s')", id, title, who)

	_, err := dbConnect.Exec(query)

	return err
}

//UsersJSON json
type UsersJSON []struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

//SendMessegeBot Call for sand messege to somebody
func SendMessegeBot(t *tgbotapi.BotAPI, who string, text string) {

	dbConnect := config.Connect()
	defer dbConnect.Close()

	var (
		pool  string
		users UsersJSON
	)
	todo := `SELECT nami.fn_telegram_sel(?)`
	sql := dbConnect.QueryRow(todo, who)
	sql.Scan(&pool)
	pool = strings.Replace(pool, `\`, ``, 1)
	err := json.Unmarshal([]byte(pool), &users)

	if err != nil {
		panic(err.Error())
	}

	for _, user := range users {

		msg := tgbotapi.NewMessage(int64(user.ID), text)
		t.Send(msg)

	}

}
