package main

import (
	"fmt"

	"github.com/romka008/calendarOfEvents/calendar"
	"github.com/romka008/calendarOfEvents/cmd"
	"github.com/romka008/calendarOfEvents/loggerNew"
	"github.com/romka008/calendarOfEvents/storage"
)

func main() {
	s := storage.NewJsonStorage("storage.json")
	// zs := storage.NewZipStorage("storage.zip")
	c := calendar.NewCalendar(s)
	logger, err := loggerNew.NewLogger("app.log")
	if err != nil {
		fmt.Println("ошибка создания файла: ", err)
	}
	cli := cmd.NewCmd(c, logger)
	cli.Run()
	defer func() {
		err := c.Save()
		close(c.Notification)
		logger.File.Close()
		if err != nil {
			fmt.Println("ошибка сериализации: ", err)
		}
	}()

	// event1, err1 := c.AddEvent("Первое событие в календаре", time.Now().Local().String(), "high")
	// if err1 != nil {
	// 	fmt.Println("Ошибка создания события:", err1)
	// 	return
	// }

	// err2 := c.SetEventReminder(event1.ID, "Сообщение о напоминании", "2025/08/14 19:16")
	// if err2 != nil {
	// 	fmt.Println("Ошибка при создании напоминания:", err2)
	// 	return
	// }
	// c.ShowEvents()
	// eer3 := c.CancelEventReminder(event1.ID)
	// if eer3 != nil {
	// 	fmt.Println("Ошибка при удалении напоминания:", eer3)
	// 	return
	// }
	// event2, err2 := c.AddEvent("Встреча", "2025/06/12 16:33", "low")
	// if err2 != nil {
	// 	fmt.Println("Ошибка создания события:", err2)
	// 	return
	// }
	// c.ShowEvents()

	// err3 := c.EditEvent(event2.ID, "Созвон", "2040/08/12 16:50", "high")
	// if err3 != nil {
	// 	fmt.Println("Ошибка при изменении события:", err3)
	// 	return
	// }
	// c.ShowEvents()

	// err4 := c.DeleteEvent(event2.ID)
	// if err4 != nil {
	// 	fmt.Println("Ошибка при удалении события:", err4)
	// 	return
	// }
	// c.ShowEvents()
}
