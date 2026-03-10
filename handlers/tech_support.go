package handlers

import (
	"context"
	"fmt"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

// Обработка нажатия кнопки "Техническая часть"
func HandleTechSupport(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCallbackUpdate) {
	// Устанавливаем состояние ожидания технической проблемы для этого чата
	waitingMutex.Lock()
	waitingForTechSupport[upd.Message.Recipient.ChatId] = true
	waitingMutex.Unlock()

	msg := maxbot.NewMessage().
		SetChat(upd.Message.Recipient.ChatId).
		SetText("Опишите вашу техническую проблему. Сообщение будет отправлено администраторам.")
	api.Messages.Send(ctx, msg)

	fmt.Printf("Tech support requested from chat ID: %d\n", upd.Message.Recipient.ChatId)
}
