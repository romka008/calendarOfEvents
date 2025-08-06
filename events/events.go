package events // имя пакета событий

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
)

type Event struct { // тип События
	ID      string    // id в формате uuid
	Title   string    // заголовок
	StartAt time.Time // дата начала
}

func (e *Event) Update(title string, date string) error {

	t, err := dateparse.ParseAny(date)

	if err != nil {
		return errors.New("неверный формат даты")
	}
	if !IsValidTitle(title) {
		return errors.New("заголовок события не соответствует условиям (длина: 3-50)")
	}

	e.Title = title
	e.StartAt = t
	return nil
}
