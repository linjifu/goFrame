package Tools

import (
	"github.com/spf13/cobra"
	"goFrame/app/Console/Commands"
)

type Command struct {
	cobra *cobra.Command
}

func NewCommand(cobra *cobra.Command) *Command {
	return &Command{
		cobra,
	}
}

func (c *Command) Register() {
	c.cobra.AddCommand(Commands.NewTestCommand().Cmd())
}
