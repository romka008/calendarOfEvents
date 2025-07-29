package main

import (
	"fmt"
	"time"

	"github.com/romka008/app/calendar"
	"github.com/romka008/app/events"
)

func main() {
	e := events.Event{ // создаем событие типа Event
		Title:   "Встреча",
		StartAt: time.Now(), // простая работа с датой :)
	}
	calendar.AddEvent("event1", e)    // добавляем событие в календарь
	fmt.Println("Календарь обновлён") // сообщаем о результате
	// fmt.Println(calendar.eventsMap)   // ошибка: Cannot use the unexported variable
}
