package scheduler

import (
	"fmt"
	"gameapp/param"
	"gameapp/service/matchingservice"
	"github.com/go-co-op/gocron"
	"sync"
	"time"
)

type Config struct {
	MatchWaitedUsersIntervalInSeconds int `koanf:"match_waited_users_interval_in_seconds"`
}

type Scheduler struct {
	sch      *gocron.Scheduler
	matchSvc matchingservice.Service
	config   Config
}

func New(matchSvc matchingservice.Service) Scheduler {
	return Scheduler{
		matchSvc: matchSvc,
		sch:      gocron.NewScheduler(time.UTC)}
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	s.sch.Every(3).Second().Do(s.MatchWaitedUsers)

	s.sch.StartAsync()

	<-done
	fmt.Println("stop scheduler..")
	s.sch.Stop()
}

func (s Scheduler) MatchWaitedUsers() {

	s.matchSvc.MatchWaitedUsers(param.MatchWaitedUsersRequest{})
}
