package calendar

import (
	"errors"
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/romka008/app/events"
)

var calendarEvents = make(map[string]events.Event) // мапа событий

func AddEvent(title string, date string) (events.Event, error) { // принимаем события в аргументе
	e, err := events.NewEvent(title, date)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
	}

	calendarEvents[e.ID] = e                         // добавляем событие используя уникальный ID как ключ
	fmt.Printf("Добавлено событие '%s' \n", e.Title) // сообщаем о результате
	return e, nil
}

func DeleteEvent(ID string) { // принимаем события в аргументе
	_, ok := calendarEvents[ID]
	if ok {
		delete(calendarEvents, ID)
		fmt.Println("Событие", ID, " удалено") // выводим лог об удалении

	} else {
		fmt.Println("Событие", ID, "не найдено в календаре")
	}
}

func EditEvent(key string, title string, dateStr string) (events.Event, error) { // принимаем события в аргументе
	_, ok := calendarEvents[key]
	t, err := dateparse.ParseAny(dateStr)

	if err != nil {
		return events.Event{}, errors.New("неверный формат даты")
	}
	if !events.IsValidTitle(title) {
		return events.Event{}, errors.New("заголовок события не соответствует условиям (длина: 3-50)")
	}
	if ok {
		event := events.Event{ID: calendarEvents[key].ID, Title: title, StartAt: t}
		calendarEvents[key] = event
		fmt.Println("Событие", key, "изменено") // выводим лог об удалении
		return event, nil

	} else {
		return events.Event{}, fmt.Errorf("событие %s не найдено в календаре", key)
	}
}

func ShowEvents() {
	for _, v := range calendarEvents {
		fmt.Println(v.Title+" - ", v.StartAt.Format("2006/01/02"))
	}
}
