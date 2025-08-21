package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
	"github.com/romka008/app/priority"
	"github.com/romka008/app/reminder"
)

type Event struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	StartAt  time.Time `json:"start_at"`
	Priority priority.Priority
	Reminder *reminder.Reminder
}

func (e *Event) Update(title string, date string, p priority.Priority) error {

	t, err := dateparse.ParseAny(date)

	if err != nil {
		return errors.New("неверный формат даты")
	}
	if !IsValidTitle(title) {
		return errors.New("заголовок события не соответствует условиям (длина: 3-50)")
	}
	err = p.Validate()
	if err != nil {
		return err
	}

	e.Title = title
	e.StartAt = t
	e.Priority = p
	return nil
}

func (e *Event) AddReminder(message string, at string, notify func(msg string)) error {
	t, err := dateparse.ParseLocal(at)
	if err != nil {
		return fmt.Errorf("неверный формат даты '%v'", at)
	}
	if e.Reminder != nil {
		return fmt.Errorf("напоминание уже добавлено")
	}
	if t.After(e.StartAt) {
		return fmt.Errorf("время напоминания должно быть перед событием. Время события: %v", e.StartAt)
	}

	duration := time.Until(t)
	if duration < 0 {
		return errors.New("пу, пу, пу. Время этого напоминания прошло")
	}

	e.Reminder = reminder.NewReminder(message, t, notify)
	err = e.Reminder.Start(duration)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) RemoveReminder() {
	if e.Reminder != nil {
		e.Reminder.Stop()
		e.Reminder = nil
	}
}
