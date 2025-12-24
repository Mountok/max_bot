package handlers

import (
	"context"
	keyboard "first-max-bot/keyboards"
	"strings"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

// Обрабатываем приветствия
func HandleGreeting(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCreatedUpdate) bool {
	if upd.Message.Body.Text == " " {
		return false
	}
	text := strings.ToLower(upd.Message.Body.Text)
	if text == "привет" || text == "здравствуйте" {
		msg := maxbot.NewMessage().
			SetChat(upd.Message.Recipient.ChatId).
			AddKeyboard(keyboard.MainKeyboard(api)).
			SetText("Привет! Я бот колледжа 😊\nВыберите действие:")

		api.Messages.Send(ctx, msg)
		return true
	}	
	return false
}