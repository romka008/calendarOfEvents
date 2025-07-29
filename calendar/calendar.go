package calendar

import (
	"fmt"

	"github.com/romka008/app/events"
)

var eventsMap = make(map[string]events.Event) // мапа событий

func AddEvent(key string, e events.Event) { // принимаем события в аргументе
	eventsMap[key] = e                         // добавляем события по ключу
	fmt.Println("Событие добавлено:", e.Title) // выводим лог для проверки
}
