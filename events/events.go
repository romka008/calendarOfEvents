package events // имя пакета событий

import "time" // импорт встроенного пакета для даты

type Event struct { // тип События
	Title   string    // заголовок
	StartAt time.Time // дата начала
}
