package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("WWGD_KEY"),
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

	b.Handle("/github", func(m *tb.Message) {
		b.Send(m.Chat, "https://github.com/wwgdelhi")
	})

	helpMessage := ` Use one of the following commands
/twitter - to get Women Who Go Delhi Twitter link
/facebook - to get a link to Women Who Go Delhi Facebook page
/github - to get a link to Women Who Go Delhi Github page
/telegram - to get an invite link for Women Who Go Delhi Telegram Group

To contribute to|modify this bot : https://github.com/hellozee/Women-Who-Go-Delhi-Bot
	`

	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Chat, helpMessage)
	})

	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		b.Send(m.Chat, "Welcome "+m.Sender+", Please introduce yourself. ")
	})

	b.Start()

}
