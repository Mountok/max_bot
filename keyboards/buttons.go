package keyboard

import (
	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func MainKeyboard(api *maxbot.Api) *maxbot.Keyboard {
	kb := api.Messages.NewKeyboardBuilder()
	kb.AddRow().
		AddLink("Открыть расписание", schemes.POSITIVE, "https://max.ru/ggkit_timetable_bot?startapp").
		AddCallback("🪨📄✂️", schemes.NEGATIVE, "num")
	return kb
}
