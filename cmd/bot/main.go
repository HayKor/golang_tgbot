package main

import (
	"log/slog"
	"os"
	"time"

	config "github.com/HayKor/golang_tgbot/pkg/config"
	handlers "github.com/HayKor/golang_tgbot/pkg/handlers"
	tele "gopkg.in/telebot.v4"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

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

	handlers.SetUpHandlers(b)

	b.Start()

}
