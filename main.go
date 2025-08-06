package main

import (
	"fmt"
	"time"

	"github.com/romka008/app/calendar"
)

func main() {
	c := calendar.NewCalendar()

	_, err1 := c.AddEvent("Первое событие в календаре", time.Now().Local().String()) // добавляем событие в календарь
	if err1 != nil {
		fmt.Println("Ошибка создания события:", err1)
		return
	}

	event2, err2 := c.AddEvent("Встреча", "2025/06/12 16:33") // добавляем событие в календарь
	if err2 != nil {
		fmt.Println("Ошибка создания события:", err2)
		return
	}
	c.ShowEvents() // вывод всех событий

	err3 := c.EditEvent(event2.ID, "Созвон", "2040/08/12 16:50")
	if err3 != nil {
		fmt.Println("Ошибка при изменении события:", err3)
		return
	}
	c.ShowEvents() // вывод всех событий

	c.DeleteEvent(event2.ID)
	c.ShowEvents() // вывод всех событий
}
