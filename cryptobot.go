package main

import (
	"time"
	"log"

	conf "github.com/vivekkartha/cryptobot/config"
	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/vivekkartha/go-binance/binance"
	"fmt"
)

func main() {

	client := binance.New(conf.BinanceApiKey, conf.BinanceSecret)
	res, err := client.GetAllOpenOrders()
	if err != nil {
		fmt.Println(err)
	}

	for _,v := range res{
		fmt.Println(v.Symbol,"Qty:",v.OrigQty,"| Price:",v.Price, "| Type:",v.Type,v.Side)
	}
	b, err := tb.NewBot(tb.Settings{
		Token:  conf.TelegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var recipient tb.Recipient

	log.Println("Bot started")
	log.Println("Listening for messages...")

	b.Handle("/start", func(m *tb.Message) {
		recipient = m.Sender
		fmt.Println("Command")
		b.Send(m.Sender,"Hey!")
		log.Println(m.Text)
	})

	b.Start()
}
