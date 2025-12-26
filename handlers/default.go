package handlers

import (
	"context"
	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)


// Если ничего не подошло
func HandleDefault(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCreatedUpdate) {
	msg := maxbot.NewMessage().
		SetChat(upd.Message.Recipient.ChatId).
		SetText("Извините, я пока понимаю только приветствие 😊")

	api.Messages.Send(ctx, msg)
}


