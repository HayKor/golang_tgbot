package middlewares

import (
	"log/slog"

	"strconv"

	tele "gopkg.in/telebot.v4"
)

func Logger() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			slog.Debug("Handled update", "id", strconv.Itoa(c.Update().ID))
			return next(c)
		}
	}
}
