package main

import (
	"github.com/go-co-op/gocron"
	"time"
)

func startHealthCheck(lb *LoadBalancer) {
	scheduler := gocron.NewScheduler(time.UTC)
	for _, srv := range lb.Servers {
		scheduler.Every(5).Seconds().Do(srv.CheckHealth)
	}
	scheduler.StartAsync()
}