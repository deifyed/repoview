package cmd

import (
	"github.com/deifyed/repoview/cmd/enroll"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	enrollCmdOpts = enroll.EnrollOptions{
		Fs: &afero.Afero{Fs: afero.NewOsFs()},
	}
	// enrollCmd represents the enroll command 
	enrollCmd = &cobra.Command{
		Use:   "enroll",
		Short: "Add a new repository to the list of repositories to be monitored",
		Args:  cobra.ExactArgs(1),
		RunE:  enroll.RunE(&enrollCmdOpts),
	}
)

func init() {
	enrollCmd.Flags().StringVarP(&enrollCmdOpts.StoragePath, "storage", "s", "", "path to storage file")
	rootCmd.AddCommand(enrollCmd)
}
