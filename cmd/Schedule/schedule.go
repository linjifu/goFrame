package Schedule

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goFrame/app/Tools"
)

type Schedule struct {
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (s Schedule) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "schedule",
		Short: "定时计划器模块",
		Long:  `定时计划器模块的启动监听`,
		Run: func(cmd *cobra.Command, args []string) {
			s.run()
		},
	}
}

func (s *Schedule) run() {
	c := cron.New(cron.WithSeconds())

	Tools.NewSchedule(c).Register()

	c.Start()

	fmt.Println("定时计划器模块启动成功")
	select {}
}
