package util

import (
	"strconv"
	"time"
)

func FormatDuration(d time.Duration) string {
	totalSeconds := int64(d / time.Second)
	var b []byte

	b = strconv.AppendInt(b, totalSeconds/3600, 10)

	b = append(b, ':')

	minutes := totalSeconds / 60 % 60
	if minutes < 10 {
		b = append(b, '0')
	}
	b = strconv.AppendInt(b, minutes, 10)

	b = append(b, ':')

	seconds := totalSeconds % 60
	if seconds < 10 {
		b = append(b, '0')
	}
	b = strconv.AppendInt(b, seconds, 10)

	return string(b)
}
