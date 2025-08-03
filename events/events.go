package events // имя пакета событий

import "time" // импорт встроенного пакета для даты

type Event struct { // тип События
	ID      string    // id в формате uuid
	Title   string    // заголовок
	StartAt time.Time // дата начала
}
