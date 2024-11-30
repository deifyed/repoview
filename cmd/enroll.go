package cmd

import (
	"github.com/deifyed/repoview/cmd/enroll"
	"github.com/spf13/cobra"
)

var (
	enrollCmdOpts = enroll.EnrollOptions{}
	// enrollCmd represents the enroll command
	enrollCmd = &cobra.Command{
		Use:   "enroll",
		Short: "Add a new repository to the list of repositories to be monitored",
		Args:  cobra.ExactArgs(1),
		RunE:  enroll.RunE(&enrollCmdOpts),
	}
)

func init() {
	rootCmd.AddCommand(enrollCmd)
}
