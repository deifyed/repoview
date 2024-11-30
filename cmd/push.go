package cmd

import (
	"github.com/deifyed/repoview/cmd/push"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var (
	pushCmdOpts = push.Options{}
	pushCmd     = &cobra.Command{
		Use:   "push",
		Short: "A brief description of your command",
		RunE:  push.RunE(&pushCmdOpts),
	}
)

func init() {
	rootCmd.AddCommand(pushCmd)
}
