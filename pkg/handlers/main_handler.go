package handlers

import (
	tele "gopkg.in/telebot.v4"
)

func SetUpHandlers(b *tele.Bot) {
	SetUpEchoHandler(b)
}
