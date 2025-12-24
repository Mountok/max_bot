package handlers

import (
	"context"
	"math/rand"
	"strconv"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

// Обрабатываем приветствия
func NUMHello(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCallbackUpdate) {
	msg := maxbot.NewMessage().
		SetChat(upd.Message.Recipient.ChatId).
		SetText("Ты хочешь поиграть в отгадай число? Иди на пары 😄")
	api.Messages.Send(ctx, msg)
}


func NUMChoice(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCallbackUpdate) bool {
	// Здесь будет логика игры в отгадай число
	num :=  rand.Intn(10)
	msg := maxbot.NewMessage().
		SetChat(upd.Message.Recipient.ChatId).
		SetText("Я загадал число от 0 до 9. Попробуй угадать!")
	api.Messages.Send(ctx, msg)

	for upd := range api.GetUpdates(ctx) {
		switch u := upd.(type) {
		case *schemes.MessageCreatedUpdate:
			guess := u.Message.Body.Text
			guessNum, _ := strconv.Atoi(guess)
			if guessNum == num {
				msg := maxbot.NewMessage().
					SetChat(u.Message.Recipient.ChatId).
					SetText("Поздравляю! Ты угадал число!")
				api.Messages.Send(ctx, msg)
				return true
			} else {
				msg := maxbot.NewMessage().
					SetChat(u.Message.Recipient.ChatId).
					SetText("Неправильно. Попробуй еще раз.")
				api.Messages.Send(ctx, msg)
			}
		}
	}
	return false
}