package calendar

import (
	"fmt"

	"github.com/romka008/app/events"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
}

func NewCalendar() *Calendar {

	var calendar = Calendar{calendarEvents: make(map[string]*events.Event)}
	return &calendar

}

func (c *Calendar) AddEvent(title string, date string) (*events.Event, error) { // принимаем события в аргументе
	e, err := events.NewEvent(title, date)
	if err != nil {
		return &events.Event{}, err
	}

	c.calendarEvents[e.ID] = e                       // добавляем событие используя уникальный ID как ключ
	fmt.Printf("Добавлено событие '%s' \n", e.Title) // сообщаем о результате
	return e, nil
}

func (c *Calendar) DeleteEvent(ID string) { // принимаем события в аргументе
	_, exists := c.calendarEvents[ID]
	if exists {
		delete(c.calendarEvents, ID)
		fmt.Println("Событие", ID, "удалено") // выводим лог об удалении
		return
	}
	fmt.Println("Событие", ID, "не найдено в календаре")

}

func (c *Calendar) EditEvent(id string, title string, dateStr string) error { // принимаем события в аргументе
	e, exists := c.calendarEvents[id]
	if !exists {
		return fmt.Errorf("событие с id='%s' не найдено в календаре", id)
	}

	err := e.Update(title, dateStr) // обновляем e через его метод
	if err != nil {
		return err
	}
	fmt.Printf("Событие с id='%s' изменено \n", id)

	return nil
}

func (c *Calendar) ShowEvents() {
	for _, v := range c.calendarEvents {
		fmt.Println(v.Title+" - ", v.StartAt.Format("2006/01/02"))
	}
}
