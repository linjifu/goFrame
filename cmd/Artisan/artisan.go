package Artisan

import (
	"fmt"
	"github.com/spf13/cobra"
	"goFrame/app/Tools"
)

type Artisan struct {
}

func NewArtisan() *Artisan {
	return &Artisan{}
}

func (s Artisan) Cmd() *cobra.Command {
	var cobra = &cobra.Command{
		Use:   "artisan",
		Short: "自定义artisan命令模块",
		Long:  `自定义artisan命令模块的启动监听`,
		Run: func(cmd *cobra.Command, args []string) {
			s.run()
		},
	}
	Tools.NewCommand(cobra).Register()

	return cobra
}

func (s *Artisan) run() {
	fmt.Println("自定义artisan命令模块启动成功")
}
