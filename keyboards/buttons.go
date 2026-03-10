package keyboard

import (
	"first-max-bot/config"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

func MainKeyboard(api *maxbot.Api, chatId int64) *maxbot.Keyboard {
	kb := api.Messages.NewKeyboardBuilder()
	kb.AddRow().
		AddLink("Открыть приложение", schemes.POSITIVE, "https://max.ru/ggkit_timetable_bot?startapp").
		AddCallback("Посмотреть расписание", schemes.POSITIVE, "getshedule")

	// Проверяем, разрешен ли этот чат для кнопки техподдержки
	for _, allowedChat := range config.AllowedTechSupportChats {
		if allowedChat == chatId {
			kb.AddRow().
				AddCallback("Техническая часть", schemes.NEGATIVE, "tech_support")
			break
		}
	}

	return kb
}
