package gigasecond

import (
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Millisecond * 1000000000000)
}
