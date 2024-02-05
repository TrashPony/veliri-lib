package utils

import "time"

func GetNowUTCTime() (int, int, int) {
	t := time.Now().In(time.UTC)

	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	return hour, minute, second
}
