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
	// delistCmd represents the delist command
	delistCmd = &cobra.Command{
		Use:   "delist",
		Short: "Remove a repository from the list of repositories to be monitored",
		Args:  cobra.ExactArgs(1),
		RunE:  delist.RunE(&delistCmdOpts),
	}
)

func init() {
	delistCmd.Flags().StringVarP(&delistCmdOpts.StoragePath, "storage", "s", "", "path to storage file")
	rootCmd.AddCommand(delistCmd)
}
