package events // имя пакета событий
import (
	"regexp"
)

func IsValidTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-Я0-9 ]{3,50}$"
	matched, err := regexp.MatchString(pattern, title) // проверка
	if err != nil {
		return false
	}
	return matched
}
