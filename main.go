package main

import (
	"context"
	"first-max-bot/config"
	"first-max-bot/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func main() {

	_, err := handlers.InitSchedule()
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	api, _ := maxbot.New(config.BotToken)

	for upd := range api.GetUpdates(ctx) {
		switch u := upd.(type) {

		case *schemes.MessageCreatedUpdate:
			// switch u.GetCommand() {
			// case "/help":
			// 	out := "команда: " + u.GetCommand()
			// 	err := api.Messages.Send(ctx, maxbot.NewMessage().SetChat(u.Message.Recipient.ChatId).SetText(out))
			// 	fmt.Printf("ansfer  %#v\n", err)
			// 	continue
			// case "/chats_full":
			// 	chatlist, err := api.Chats.GetChats(ctx, 0, 0)
			// 	if err != nil {
			// 		log.Printf("Unknown type: %#v\n", err)
			// 	}
			// 	out := fmt.Sprintf("AllChats:\n")
			// 	for _, chat := range chatlist.Chats {
			// 		out += fmt.Sprintf("title: %#v\n", chat.Title)
			// 		out += fmt.Sprintf("id: %#v\n", chat.ChatId)
			// 		out += fmt.Sprintf("description: %#v\n", chat.Description)
			// 		out += fmt.Sprintf("is public: %#v\n", chat.IsPublic)
			// 		out += fmt.Sprintf("link: %#v\n", chat.Link)
			// 		out += fmt.Sprintf("status: %#v\n", chat.Status)
			// 		out += fmt.Sprintf("owner: %#v\n", chat.OwnerId)
			// 		out += fmt.Sprintf("type: %#v\n", chat.Type)
			// 		out += "______\n"
			// 	}
			// 	api.Messages.Send(ctx, maxbot.NewMessage().SetChat(u.Message.Recipient.ChatId).SetText(out))
			// 	continue
			// case "/sm":
			// 	chatlist, err := api.Chats.GetChats(ctx, 0, 0)
			// 	if err != nil {
			// 		log.Printf("Unknown type: %#v\n", err)
			// 	}
			// 	for _, chat := range chatlist.Chats {
			// 		fmt.Sprintf("id: %#v\n", chat.ChatId)
			// 		api.Messages.Send(ctx, maxbot.NewMessage().SetChat(chat.ChatId).SetText("ЙООООООООВ"))		
			// 	}
			// 	continue
			// }

			// keyboard := api.Messages.NewKeyboardBuilder()
			// keyboard.
			// 	AddRow().
			// 	AddGeolocation("Прислать геолокацию", true).
			// 	AddContact("Прислать контакт")
			// keyboard.
			// 	AddRow().
			// 	AddLink("Ссылка", schemes.POSITIVE, "https://max.ru").
			// 	AddCallback("Аудио", schemes.NEGATIVE, "audio").
			// 	AddCallback("Видео", schemes.NEGATIVE, "video")
			// keyboard.
			// 	AddRow().
			// 	AddCallback("Картинка", schemes.POSITIVE, "picture")
			
			
			// api.Messages.Send(ctx, maxbot.NewMessage().SetChat(u.Message.Recipient.ChatId).AddKeyboard(keyboard).SetText("выбери"))
			// api.Messages.Send(ctx, maxbot.NewMessage().Reply("**Reply** universal", u.Message).SetFormat("markdown"))

			if !handlers.HandleGreeting(ctx, api, u) {
				handlers.HandleDefault(ctx, api, u)
			}
		case *schemes.MessageCallbackUpdate:
			if u.Callback.Payload == "getshedule" {
				handlers.HandleSchedule(ctx, api, u)
			}
		
		default:
			log.Printf("Unknown type: %#v", upd)
		}
	}
}
