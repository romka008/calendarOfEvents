package events // имя пакета событий
import (
	"errors"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

func NewEvent(title string, dateStr string) (*Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return &Event{}, errors.New("неверный формат даты")
	}
	if !IsValidTitle(title) {
		return &Event{}, errors.New("заголовок события не соответствует условиям (длина: 3-50)")
	}

	return &Event{
		ID:      uuid.NewString(),
		Title:   title,
		StartAt: t,
	}, nil
}
