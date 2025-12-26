package main

import (
	"context"
	"first-max-bot/config"
	"first-max-bot/handlers"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func main() {

	schedule, err := handlers.InitSchedule()
	if err != nil {
		panic(err)
	}
	fmt.Println(schedule.Groups[0].Name)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()



	api, _ := maxbot.New(config.BotToken)

	for upd := range api.GetUpdates(ctx) {
		switch u := upd.(type) {
		case *schemes.MessageCreatedUpdate:
			if !handlers.HandleGreeting(ctx, api, u) {
				handlers.HandleDefault(ctx, api, u)
			}
		}
	}
}
