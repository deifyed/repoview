package cmd

import (
	"github.com/deifyed/repoview/cmd/status"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var (
	statusCmdOpts = status.Options{
		Fs: &afero.Afero{Fs: afero.NewOsFs()},
	}
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Show repository status for all enrolled repositories",
		RunE:  status.RunE(&statusCmdOpts),
	}
)

func init() {
	statusCmd.PreRun = func(cmd *cobra.Command, args []string) {
		statusCmdOpts.StoragePath = viper.GetString("storage.path")
	}

	rootCmd.AddCommand(statusCmd)
}
