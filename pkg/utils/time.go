package utils

import "time"

type (
	Clock interface {
		Now() time.Time
	}
)

var (
	CurrentTime Clock
)

func GetCurrentTime() time.Time {
	return CurrentTime.Now()
}
