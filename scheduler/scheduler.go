package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct {
}

func New() Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start(done chan bool) {
	fmt.Println("scheduler started")
	for {
		select {
		case <-done:
			fmt.Println("exitttting....")
			return
		default:
			now := time.Now()
			fmt.Println("scheduler now", now)
			time.Sleep(3 * time.Second)
		}
	}
}
