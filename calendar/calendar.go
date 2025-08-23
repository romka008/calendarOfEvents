package calendar

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/romka008/calendarOfEvents/events"
	"github.com/romka008/calendarOfEvents/priority"
	"github.com/romka008/calendarOfEvents/storage"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
	Notification   chan string
}

func NewCalendar(s storage.Store) *Calendar {
	var calendar = Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        s,
		Notification:   make(chan string)}
	return &calendar
}

func (c *Calendar) Save() error {

	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return err
	}
	c.Notification <- ("JSON строка: " + string(data))
	return c.storage.Save(data)
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return fmt.Errorf("Ошибка десериализации: %s", err)
	}
	return err
}

func (c *Calendar) AddEvent(title string, date string, priority priority.Priority) (*events.Event, error) {
	e, err := events.NewEvent(title, date, priority)
	if err != nil {
		return nil, err
	}

	c.calendarEvents[e.ID] = e
	return e, nil
}

func (c *Calendar) DeleteEvent(id string) error {
	_, err := c.checkEventExist(id)
	if err != nil {
		return err
	}
	delete(c.calendarEvents, id)
	return nil

}

func (c *Calendar) EditEvent(id string, title string, dateStr string, priority priority.Priority) error {
	e, err := c.checkEventExist(id)
	if err != nil {
		return err
	}

	err = e.Update(title, dateStr, priority)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calendar) ShowEvents() {
	for _, v := range c.calendarEvents {
		c.Notification <- (v.Title + " - " + v.StartAt.Format("2006/01/02 15:04") + " - " + v.ID)
	}
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) SetEventReminder(id string, message string, at string) error {
	e, err := c.checkEventExist(id)
	if err != nil {
		return err
	}
	return e.AddReminder(message, at, c.Notify)
}

func (c *Calendar) CancelEventReminder(id string) error {
	e, err := c.checkEventExist(id)
	if err != nil {
		return err
	}
	if e.Reminder == nil {
		return errors.New("у события отсутствует напоминание")
	}
	e.RemoveReminder()
	return nil
}
func (c *Calendar) checkEventExist(id string) (*events.Event, error) {
	e, exists := c.calendarEvents[id]
	if !exists {
		return nil, errors.New("событие не найдено в календаре c id " + id)
	}
	return e, nil
}
