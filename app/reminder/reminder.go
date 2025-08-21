package reminder

import (
	"time"
)

type Reminder struct {
	Message  string
	At       time.Time
	Sent     bool
	timer    *time.Timer
	notifier func(msg string)
}

func (r *Reminder) Start(d time.Duration) error {

	r.timer = time.AfterFunc(d, r.Send)
	r.notifier("Напоминание произойдет через " + d.String())
	return nil
}

func NewReminder(message string, at time.Time, callback func(msg string)) *Reminder {

	return &Reminder{
		Message:  message,
		At:       at,
		Sent:     false,
		timer:    nil,
		notifier: callback,
	}
}

func (r *Reminder) Send() {

	if r.Sent {
		return
	}
	r.notifier(r.Message)
	r.Sent = true
}

// add "Встреча" "2025/08/21 22:50" "low"
// add_reminder "61aea15b-f749-4013-b374-e48b5897c6a0" "какое-то сообщение" "2025/08/21 22:41"

func (r *Reminder) Stop() {
	if r.timer != nil {
		r.Sent = true
		r.timer = nil
	}
}
