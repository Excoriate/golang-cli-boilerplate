package utils

import "time"

func NewUnixTimeWithAddedMinutes(number int64) int64 {
	if number == 0 {
		return 0
	}

	return time.Now().Add(time.Duration(number) * time.Minute).Unix()
}
