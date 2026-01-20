package handlers

import (
	"context"
	"fmt"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

// Если ничего не подошло
func HandleDefault(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCreatedUpdate) {
	// Проверяем, ожидает ли пользователь ввод группы
	waitingMutex.RLock()
	isWaiting := waitingForGroup[upd.Message.Recipient.ChatId]
	waitingMutex.RUnlock()

	if isWaiting {
		// Пытаемся обработать как название группы
		groupName := upd.Message.Body.Text
		group := GetScheduleForGroup(globalSchedule, groupName)

		if group != nil {
			// Группа найдена, показываем расписание
			day := GetScheduleForDay(group, GetCurrentDayName())
			if day != nil && len(day.Lessons) > 0 {
				messageText := fmt.Sprintf("📆 Расписание для группы %s на сегодня (%s):\n", groupName, GetCurrentDayName())

				if len(day.Lessons) >= 1 {
					messageText += fmt.Sprintf("1️⃣ пара: %s\n", day.Lessons[0].Subject)
				}
				if len(day.Lessons) >= 3 {
					messageText += fmt.Sprintf("2️⃣ пара: %s\n", day.Lessons[2].Subject)
				}
				if len(day.Lessons) >= 5 {
					messageText += fmt.Sprintf("3️⃣ пара: %s\n", day.Lessons[4].Subject)
				}

				msg := maxbot.NewMessage().
					SetChat(upd.Message.Recipient.ChatId).
					SetText(messageText)
				api.Messages.Send(ctx, msg)
			} else {
				msg := maxbot.NewMessage().
					SetChat(upd.Message.Recipient.ChatId).
					SetText(fmt.Sprintf("На сегодня (%s) занятий нет для группы %s", GetCurrentDayName(), groupName))
				api.Messages.Send(ctx, msg)
			}
		} else {
			// Группа не найдена
			msg := maxbot.NewMessage().
				SetChat(upd.Message.Recipient.ChatId).
				SetText(fmt.Sprintf("Группа '%s' не найдена. Отправьте мне правильное название вашей группы (пример: И1.22-12)", groupName))
			api.Messages.Send(ctx, msg)
			return // Не сбрасываем состояние ожидания, чтобы пользователь мог попробовать снова
		}

		// Сбрасываем состояние ожидания
		waitingMutex.Lock()
		delete(waitingForGroup, upd.Message.Recipient.ChatId)
		waitingMutex.Unlock()

		return
	}

	msg := maxbot.NewMessage().
		SetChat(upd.Message.Recipient.ChatId).
		SetText("Извините, я пока учусь 😊")

	api.Messages.Send(ctx, msg)
}

func HandleSchedule(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCallbackUpdate) {
	// Устанавливаем состояние ожидания группы для этого чата
	waitingMutex.Lock()
	waitingForGroup[upd.Message.Recipient.ChatId] = true
	waitingMutex.Unlock()

	msg := maxbot.NewMessage().
		SetChat(upd.Message.Recipient.ChatId).
		SetText(fmt.Sprintf("Отправьте мне правильное название вашей группы (пример: И1.22-12)"))
	api.Messages.Send(ctx, msg)

	fmt.Printf("Schedule requested from chat ID: %d\n", upd.Message.Recipient.ChatId)
}
