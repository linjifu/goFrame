package Tools

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"goFrame/app/Console/CronJobs"
)

type Schedule struct {
	cron *cron.Cron
}

func NewSchedule(cron *cron.Cron) *Schedule {
	return &Schedule{
		cron,
	}
}

func (s *Schedule) Register() {
	if _, err := s.cron.AddJob("*/3 * * * * *", CronJobs.NewTestCron()); err != nil {
		fmt.Println(err)
	}
}
