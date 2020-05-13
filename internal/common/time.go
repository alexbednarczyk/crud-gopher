package common

import (
	"fmt"
	"time"
)

// CurrentTime based on time.Now and RFC3339
func CurrentTime() (time.Time, error) {
	timeNow := time.Now().UTC().Format(time.RFC3339)
	formattedTime, err := time.Parse(time.RFC3339, timeNow)
	if err != nil {
		return time.Now(), fmt.Errorf("Failed to parse time. %v", err)
	}
	return formattedTime, nil
}
