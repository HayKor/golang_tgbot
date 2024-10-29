package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/HayKor/golang_tgbot/pkg/config"
	tele "gopkg.in/telebot.v4"
)

func main() {
	config.SetDotEnv()
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		slog.Error(err.Error())
		return
	} else {
		slog.Info("Started new bot.")
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()

}
