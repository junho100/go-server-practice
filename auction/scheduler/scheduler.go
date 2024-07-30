package scheduler

import (
	"auction/service"
	"log"
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
	log.Println("check status of auctions.....")

	auctions, _ := s.Service.GetExpiredAuctionByTime(time.Now())

	if len(auctions) != 0 {
		log.Println("expired auction detected.")
	}

	for _, auction := range auctions {
		s.Service.TerminateAuction(&auction)
	}
}
