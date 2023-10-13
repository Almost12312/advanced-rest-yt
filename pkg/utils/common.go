package utils

import "time"

func DoWithAttempts(fn func() error, attempts uint, delay time.Duration) (err error) {
	for attempts > 0 {
		err := fn()
		if err != nil {
			time.Sleep(delay)
			attempts--

			continue
		} else {
			return nil
		}
	}

	return
}
