package main

import (
	//"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

var b tb.Bot

func telegram(m *tb.Message) {
	b.Send(m.Chat, "https://t.me/womenwhogodelhi")
}

func facebook(m *tb.Message) {
	b.Send(m.Chat, "https://www.facebook.com/WomenWhoGoDelhi/")
}

func twitter(m *tb.Message) {
	b.Send(m.Chat, "https://twitter.com/womenwhogo_del")
}

func github(m *tb.Message) {
	b.Send(m.Chat, "https://github.com/wwgdelhi")
}

func help(m *tb.Message) {
	helpMessage := ` Use one of the following commands
/twitter - to get Women Who Go Delhi Twitter link
/facebook - to get a link to Women Who Go Delhi Facebook page
/github - to get a link to Women Who Go Delhi Github page
/telegram - to get an invite link for Women Who Go Delhi Telegram Group

To contribute to|modify this bot : https://github.com/hellozee/Women-Who-Go-Delhi-Bot
`
	b.Send(m.Chat, helpMessage)
}

func welcome(m *tb.Message) {
	message := "Welcome "

	if m.UserJoined.Username != "" {
		message += "@" + m.UserJoined.Username
	} else {
		message += m.UserJoined.FirstName
	}

	message += ", please introduce yourself"

	b.Send(m.Chat, message)
}

func main() {

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("WWGD_KEY"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/telegram", telegram)

	b.Handle("/twitter", twitter)

	b.Handle("/facebook", facebook)

	b.Handle("/github", github)

	b.Handle("/help", help)

	b.Handle(tb.OnUserJoined, welcome)

	b.Start()

}
