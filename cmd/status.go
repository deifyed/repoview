package cmd

import (
	"github.com/deifyed/repoview/cmd/status"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var (
	statusCmdOpts = status.StatusOptions{}
	statusCmd     = &cobra.Command{
		Use:   "status",
		Short: "Show repository status for all enrolled repositories",
		RunE:  status.RunE(&statusCmdOpts),
	}
)

func init() {
	rootCmd.AddCommand(pushCmd)
}
