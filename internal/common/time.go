package common

import "time"

// CurrentTime returns the current UTC time with sub-second precision removed so
// values are deterministic in tests and logs.
func CurrentTime() time.Time {
	return time.Now().UTC().Truncate(time.Millisecond)
}
