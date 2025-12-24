package main

import (
	"context"
	"first-max-bot/config"
	"first-max-bot/handlers"
	"os"
	"os/signal"
	"syscall"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	api, _ := maxbot.New(config.BotToken)

	for upd := range api.GetUpdates(ctx) {
		switch u := upd.(type) {
		case *schemes.MessageCreatedUpdate:
			if !handlers.HandleGreeting(ctx, api, u) {
				handlers.HandleDefault(ctx, api, u)
			}
		case *schemes.MessageCallbackUpdate:
			// Обработка колбэков будет здесь
			if u.Callback.Payload == "num" {
				handlers.NUMChoice(ctx, api, u)
			}
			
		}
	}
}
