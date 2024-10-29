package handlers

import (
	middlewares "github.com/HayKor/golang_tgbot/pkg/middlewares"
	tele "gopkg.in/telebot.v4"
)

func SetUpEchoHandler(b *tele.Bot) {
	g := b.Group()
	g.Use(middlewares.Logger())

	g.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(c.Message().Text)
	})

	g.Handle(tele.OnPhoto, func(c tele.Context) error {
		return c.Send(c.Message().Photo)
	})

}
