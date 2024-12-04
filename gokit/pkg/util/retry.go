package util

import (
	"fmt"
	"time"
)

type DoResult []any
type DoFunc func() (DoResult, error)

func Retry(do DoFunc, maxAttempt int, sleepBetweenAttempt *time.Duration) (DoResult, error) {
	attempt := 0
	var lastErr error
	for attempt < maxAttempt {
		results, err := do()
		if err != nil {
			lastErr = err
			attempt++

			if sleepBetweenAttempt != nil {
				time.Sleep(*sleepBetweenAttempt)
			}

			continue
		}

		return results, nil
	}

	return nil, fmt.Errorf("reach max retry, the last error was %v", lastErr)
}
