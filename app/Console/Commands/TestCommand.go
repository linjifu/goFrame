package Commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

type TestCommand struct {
}

func NewTestCommand() *TestCommand {
	return &TestCommand{}
}

func (t *TestCommand) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "artisan test 模块",
		Long:  `artisan test 模块的启动监听`,
		Run: func(cmd *cobra.Command, args []string) {
			t.run()
		},
	}
}

func (a *TestCommand) run() {
	fmt.Println("artisan test 模块 执行了")
}
