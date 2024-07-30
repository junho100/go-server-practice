package scheduler

import (
	"auction/service"
	"time"

	"github.com/robfig/cron"
)

type Scheduler struct {
	Service *service.Service
}

func (s *Scheduler) InitiateTask() {
	cronJob := cron.New()

	_ = cronJob.AddFunc("* * * * * *", s.InitiateAuctionTask)

	cronJob.Start()
}

func (s *Scheduler) InitiateAuctionTask() {
	auctions, _ := s.Service.GetExpiredAuctionByTime(time.Now())

	for _, auction := range auctions {
		s.Service.TerminateAuction(&auction)
	}
}
