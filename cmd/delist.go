package cmd

import (
	"github.com/deifyed/repoview/cmd/delist"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	delistCmdOpts = delist.DelistOptions{
		Fs: &afero.Afero{Fs: afero.NewOsFs()},
	}
	// enrollCmd represents the enroll command
	delistCmd = &cobra.Command{
		Use:   "delist",
		Short: "Remove a repository from the list of repositories to be monitored",
		Args:  cobra.ExactArgs(1),
		RunE:  delist.RunE(&delistCmdOpts),
	}
)

func init() {
	rootCmd.AddCommand(delistCmd)
}
