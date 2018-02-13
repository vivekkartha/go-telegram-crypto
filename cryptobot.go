package main

import (
	"time"
	"log"

	conf "./config"
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

	fmt.Println("All open orders")
	for _, v := range res {
		fmt.Println(v.Symbol, "Qty:", v.OrigQty, "| Price:", v.Price, "| Type:", v.Type, v.Side)
	}

	symbol := "QSPBTC"
	fmt.Printf("Trade history - %s\n", symbol)
	trades, err := client.GetTrades("QSPBTC")
	for _, t := range trades {
		side := "SELL"
		if t.IsBuyer {
			side = "BUY"
		}
		totalPrice := t.Price * t.Quantity
		fmt.Printf("%s Rate:%f  Qty:%d  Tot:%f\n Commission:%f(%s) %s\n",
			symbol, t.Price, uint64(t.Quantity), totalPrice, t.Commission, t.CommissionAsset, side)
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
		b.Send(m.Sender, "Hey!")
		log.Println(m.Text)
	})

	b.Start()
}
