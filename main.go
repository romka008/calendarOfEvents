package main

import (
	"fmt"
	"time"

	"github.com/romka008/app/calendar"
)

func main() {
	_, err1 := calendar.AddEvent("Первое событие в календаре", time.Now().Local().String()) // добавляем событие в календарь
	if err1 != nil {
		fmt.Println("Ошибка создания события:", err1)
	}

	event2, err2 := calendar.AddEvent("Встреча2", "2025/06/12 16:33") // добавляем событие в календарь
	if err2 != nil {
		fmt.Println("Ошибка создания события:", err2)
	}
	calendar.ShowEvents() // вывод всех событий

	_, err3 := calendar.EditEvent(event2.ID, "Созвон", "2040/08/12 16:50")
	if err3 != nil {
		fmt.Println("Ошибка при изменении события:", err3)
	}
	calendar.ShowEvents() // вывод всех событий

	calendar.DeleteEvent(event2.ID)
	calendar.ShowEvents() // вывод всех событий

	// e, err = events.NewEvent("Встреча", "2025/06/12 16:33")
	// if err != nil {
	// 	fmt.Println("Ошибка создания события:", err)
	// }
	// calendar.AddEvent("event2", e)

	// // calendar.DeleteEvent("event3")
	// _, err = calendar.EditEvent("event2d", "Созвон", "2040/08/12 16:50")
	// if err != nil {
	// 	fmt.Println("Ошибка при изменении события:", err)
	// }
	// // fmt.Printf("Добавлено событие '%s' \n", e.Title)
	// calendar.ShowEvents() // вывод всех событий
}
