package utils

import "time"

// Schedule ...
type Schedule interface {
	Run()
}

type schedule struct {
	Schedule
}

func (s *schedule) Run(runner func(), interval time.Duration) {
	ticker := time.NewTicker(1 * time.Second)
	go func(runFunc func()) {

		for range ticker.C {

		}

	}(runner)

	// wait for 10 seconds
	time.Sleep(interval)
	ticker.Stop()
}
