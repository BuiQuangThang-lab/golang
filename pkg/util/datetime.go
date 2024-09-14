package util

import (
	"time"
)

func FormatDateTime(t time.Time, outputLayout string) string {
	return t.Format(outputLayout)
}
