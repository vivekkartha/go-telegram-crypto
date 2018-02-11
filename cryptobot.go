package main

import (
	"time"
	"log"
	tb "gopkg.in/tucnak/telebot.v2"

	conf "github.com/vivekkartha/cryptobot/config"
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

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hey there!")
	})
	b.Start()
}
