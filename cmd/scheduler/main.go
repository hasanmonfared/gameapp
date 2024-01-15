package main

import (
	"gameapp/config"
	"gameapp/scheduler"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("config.yml")
	done := make(chan bool)
	go func() {
		sch := scheduler.New()
		sch.Start(done)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	done <- true
	time.Sleep(cfg.Application.GracefulShutdownTimeout)
}
