package main

import (
	"time"
	"log"


	conf "github.com/vivekkartha/cryptobot/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	b, err := tb.NewBot(tb.Settings{
		Token:  conf.TelegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Bot started")
	log.Println("Listening for messages...")
	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hey there!")
	})
	b.Start()
}
