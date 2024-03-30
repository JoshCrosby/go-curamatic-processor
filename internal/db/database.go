package main

import (
	"time"
)

func scheduleTask(interval time.Duration, task func()) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				task()
			}
		}
	}()
}

// Example usage within main.go to run a task weekly
// scheduleTask(168*time.Hour, myTaskFunction)
