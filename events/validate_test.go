package events

import (
	"fmt"
	"testing"
)

func TestIsValidTitle(t *testing.T) {
	result := IsValidTitle("123456789")
	fmt.Println(result)

	if !result {
		t.Errorf("Должно вернуться true, вернулось %v", result)
	}

	result = IsValidTitle("123")
	if !result {
		t.Errorf("Длина заголовка = минимальной длине. Должно вернуться true, вернулось %v", result)
	}

	result = IsValidTitle("12")
	fmt.Println(result)
	if result {
		t.Errorf("Вернулось %v, должно вернуться false", result)
	}

	result = IsValidTitle("")
	fmt.Println(result)
	if result {
		t.Errorf("Вернулось %v, должно вернуться false", result)
	}
}
