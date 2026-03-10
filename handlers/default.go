package handlers

import (
	"context"
	"fmt"

	"first-max-bot/config"

	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

// Если ничего не подошло
func HandleDefault(ctx context.Context, api *maxbot.Api, upd *schemes.MessageCreatedUpdate) {
	// Проверяем, ожидает ли пользователь ввод технической проблемы
	waitingMutex.RLock()
	isWaitingTech := waitingForTechSupport[upd.Message.Recipient.ChatId]
	waitingMutex.RUnlock()

	if isWaitingTech {
		// Отправляем сообщение всем администраторам
		techMessage := fmt.Sprintf("🆘 Техническая проблема от пользователя %d:\n\n%s",
			upd.Message.Recipient.ChatId, upd.Message.Body.Text)

		for _, adminID := range config.AdminUserIDs {
			adminMsg := maxbot.NewMessage().
				SetChat(adminID).
				SetText(techMessage)
			api.Messages.Send(ctx, adminMsg)
		}

		// Отправляем подтверждение пользователю
		confirmMsg := maxbot.NewMessage().
			SetChat(upd.Message.Recipient.ChatId).
			SetText("Спасибо! Ваше сообщение отправлено администраторам. Мы свяжемся с вами в ближайшее время.")
		api.Messages.Send(ctx, confirmMsg)

		// Сбрасываем состояние ожидания
		waitingMutex.Lock()
		delete(waitingForTechSupport, upd.Message.Recipient.ChatId)
		waitingMutex.Unlock()

		return
	}

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

	// msg := maxbot.NewMessage().
	// 	SetChat(upd.Message.Recipient.ChatId).
	// 	SetText("Извините, я пока учусь 😊")

	// api.Messages.Send(ctx, msg)
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
