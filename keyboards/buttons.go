package keyboard

import (
	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func MainKeyboard(api *maxbot.Api) *maxbot.Keyboard {
	kb := api.Messages.NewKeyboardBuilder()
	kb.AddRow().
		AddLink("Открыть приложение", schemes.POSITIVE, "https://max.ru/ggkit_timetable_bot?startapp").
		AddCallback("Посмотреть расписание", schemes.POSITIVE, "getshedule")
	return kb
}
