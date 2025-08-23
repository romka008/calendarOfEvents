package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/romka008/calendarOfEvents/calendar"
	"github.com/romka008/calendarOfEvents/loggerNew"
	"github.com/romka008/calendarOfEvents/priority"
)

type Cmd struct {
	calendar  *calendar.Calendar
	logger    *Logger
	newLogger *loggerNew.Logger
}

func NewCmd(c *calendar.Calendar, l *loggerNew.Logger) *Cmd {
	return &Cmd{
		calendar:  c,
		logger:    &Logger{},
		newLogger: l,
	}
}

func (c *Cmd) executor(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}
	c.logger.AddLog("> " + input)

	parts, _ := shlex.Split(input)
	cmd := strings.ToLower(parts[0])
	switch cmd {

	case "add":
		if len(parts) < 4 {
			c.Print("Формат: add \"название события\" \"дата события\" \"приоритет\"", "Error")
			return
		}

		title := parts[1]
		date := parts[2]
		priority := priority.Priority(parts[3])

		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			c.Print("Ошибка добавления:"+err.Error(), "Error")
		} else {
			c.Print("Событие: '"+e.Title+"' добавлено ID: '"+e.ID+"'", "Info")
		}
	case "add_reminder":
		if len(parts) < 3 {
			c.Print("Формат: add_reminder \"id события\" \"сообщение напоминания\" \"дата и время напоминания\"", "Info")
			return
		}

		id := parts[1]
		message := parts[2]
		date := parts[3]

		err := c.calendar.SetEventReminder(id, message, date)
		if err != nil {
			c.Print("Ошибка добавления напоминания: "+err.Error(), "Error")
		}
	case "cancel_reminder":
		if len(parts) < 2 {
			c.Print("Формат: cancel_reminder \"id события\"", "Info")
			return
		}

		id := parts[1]

		err := c.calendar.CancelEventReminder(id)
		if err != nil {
			c.Print("Ошибка остановки и удаления напоминания:"+err.Error(), "Error")
		} else {
			c.Print("Напоминание удалено", "Info")
		}
	case "update":
		if len(parts) < 5 {
			c.Print("Формат: update \"id события\" \"название события\" \"дата события\" \"приоритет\"", "Info")
			return
		}

		id := parts[1]
		title := parts[2]
		date := parts[3]
		priority := priority.Priority(parts[4])

		err := c.calendar.EditEvent(id, title, date, priority)
		if err != nil {
			c.Print("Ошибка при изменении:"+err.Error(), "Error")
		} else {
			fmt.Printf("Событие:\"%s\" изменено\n", title)
		}
	case "remove":
		if len(parts) < 1 {
			c.Print("Формат: remove \"id события\"", "Info")
			return
		}

		id := parts[1]

		err := c.calendar.DeleteEvent(id)
		if err != nil {
			c.Print("Ошибка добавления:"+err.Error(), "Error")
		} else {
			fmt.Printf("Событие: \"%s\" удалено\n", id)
		}
	case "log":
		logs := c.logger.GetLog()
		for _, log := range logs {
			c.Print("time:"+log.At+"value:"+log.Value, "Info")
		}
	case "list":
		c.calendar.ShowEvents()
	case "help":
		c.Print("В программе доступны следующие команды:", "Info")
		c.Print("Добавления события: 'add \"название события\" \"дата события\" \"приоритет\"'", "Info")
		c.Print("Редактирование события: 'update \"id события\" \"название события\" \"дата события\" \"приоритет\"'", "Info")
		c.Print("Удаление события: 'remove \"id события\"'", "Info")
		c.Print("Добавления напоминания для события: 'add_reminder \"id события\" \"сообщение для напоминания\" \"дата и время напоминания\"'", "Info")
		c.Print("Остановка и удаление напоминания для события: 'cancel_reminder \"id события\"'", "Info")
		c.Print("Отобразить все события: 'list'", "Info")
		c.Print("Выход из программы: 'exit'", "Info")
		c.Print("Вывод логов: 'log'", "Info")

	case "exit":
		err := c.calendar.Save()
		if err != nil {
			c.Print(err.Error(), "Error")
			return
		}
		os.Exit(0)

	default:
		c.Print("Неизвестная команда:", "Info")
		c.Print("Введите 'help' для списка команд", "Info")
	}
}
func (*Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Добавить событие"},
		{Text: "add_reminder", Description: "Добавить напоминание для события"},
		{Text: "cancel_reminder", Description: "Остановить и удалить напоминание для события"},
		{Text: "list", Description: "Показать все события"},
		{Text: "remove", Description: "Удалить событие"},
		{Text: "update", Description: "Обновить событие"},
		{Text: "help", Description: "Показать справку"},
		{Text: "exit", Description: "Выйти из программы"},
		{Text: "log", Description: "Вывод всех логов в консоль"},
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	go func() {
		for msg := range c.calendar.Notification {
			c.Print(msg, "")
		}
	}()
	p.Run()
}

func (c *Cmd) Print(text string, l loggerNew.LoggerType) {
	fmt.Println(text)
	c.logger.AddLog("> " + text)
	if loggerNew.InfoLoggerType == l {
		c.newLogger.Info(text)
	} else if loggerNew.ErrorLoggerType == l {
		c.newLogger.Error(text)
	}
}
