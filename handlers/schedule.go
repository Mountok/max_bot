package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// Глобальная мапа для отслеживания пользователей, ожидающих ввод группы
var waitingForGroup = make(map[int64]bool)

// Глобальная мапа для отслеживания пользователей, ожидающих ввод технической проблемы
var waitingForTechSupport = make(map[int64]bool)
var waitingMutex sync.RWMutex

// Глобальная переменная для расписания
var globalSchedule *Schedule

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

	globalSchedule = &schedule
	return &schedule, nil
}

func GetScheduleForGroup(schedule *Schedule, groupName string) *Group {
	for _, group := range schedule.Groups {
		if group.Name == groupName {
			return &group
		}
	}
	return nil
}

func GetScheduleForDay(group *Group, dayName string) *Day {
	for _, days := range group.Days {
		if days.Name == dayName {
			return &days
		}
	}
	return nil
}

// Функция для получения текущего дня недели
func GetCurrentDayName() string {
	days := []string{"Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"}
	weekday := time.Now().Weekday()
	return days[weekday]
}
