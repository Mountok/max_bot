package handlers


import (
	"encoding/json"
	"fmt"
	"os"
)

// Структура для занятия
type Lesson struct {
	Subject   string `json:"subject"`   // Предмет
	Teacher   string `json:"teacher"`   // Преподаватель
	Room      string `json:"room"`      // Кабинет
	Time      string `json:"time"`      // Время занятия
	LessonNum int    `json:"lessonNum"` // Номер пары
}

// Структура для дня недели
type Day struct {
	Name    string   `json:"name"`    // Название дня (Понедельник, Вторник и т.д.)
	Lessons []Lesson `json:"lessons"` // Занятия в этот день
}

// Структура для группы
type Group struct {
	Name string `json:"name"` // Название группы
	Days []Day  `json:"days"` // Дни недели с занятиями
}

// Структура для всего расписания
type Schedule struct {
	Groups []Group `json:"groups"` // Все группы
}

func InitSchedule() (*Schedule, error) {
	
	data, err := os.ReadFile("schedule.json")
	if err != nil {
		panic(err)
	}

	var schedule Schedule

	err = json.Unmarshal(data, &schedule)
	if err != nil {
		panic(err)
	}

	fmt.Println("ParseData")
	
	return &schedule, nil
}
