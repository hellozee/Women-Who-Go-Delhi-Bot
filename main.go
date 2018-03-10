package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  WWGDel_Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/telegram", func(m *tb.Message) {
		b.Send(m.Chat, "https://t.me/womenwhogodelhi")
		fmt.Println("Got a request")
	})

	b.Handle("/twitter", func(m *tb.Message) {
		b.Send(m.Chat, "https://twitter.com/womenwhogo_del")
	})

	b.Handle("/facebook", func(m *tb.Message) {
		b.Send(m.Chat, "https://www.facebook.com/WomenWhoGoDelhi/")
	})

	b.Start()

}
