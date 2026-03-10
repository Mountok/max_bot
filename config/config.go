package config

var BotToken = "f9LHodD0cOIBuENhVtNNEwxSMzzETMN9D0_a0nnCH4Aqwdbu2g3exk2uBOqFZQLGmY3kPuwZ80doJiye3UpI" // можно хранить токен в переменной окружения

// Список ID администраторов для получения технических проблем
var AdminUserIDs = []int64{
	2804823,
	// Добавьте сюда ID администраторов
	// Пример: 123456789,
}

// Список разрешенных ChatId для показа кнопки "Техническая часть"
var AllowedTechSupportChats = []int64{
	-69370301647575,
	// Добавьте сюда ID чатов, где должна отображаться кнопка "Техническая часть"
	// Пример: 123456789,
}
