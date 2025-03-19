package handbook

import (
	"fmt"
	"testing"
)

func TestGetYandexToken(t *testing.T) {
	token, err := getYandexToken("")
	if err != nil {
		return
	}

	fmt.Println(token)

	data, err := TranslateText(token, "", "en", []string{"Привет, мир!"}) // "en", "zh"
	if err != nil {
		return
	}

	fmt.Println(data.Translations[0].Text)
}
