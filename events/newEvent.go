package events

import (
	"errors"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/romka008/calendarOfEvents/priority"
)

func NewEvent(title string, dateStr string, priority priority.Priority) (*Event, error) {
	t, err := dateparse.ParseLocal(dateStr)
	if err != nil {
		return nil, errors.New("неверный формат даты")
	}
	if !IsValidTitle(title) {
		return nil, errors.New("заголовок события не соответствует условиям (длина: 3-50)")
	}
	err = priority.Validate()
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:       uuid.NewString(),
		Title:    title,
		StartAt:  t,
		Priority: priority,
		Reminder: nil,
	}, nil
}
