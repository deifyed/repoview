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
		Args:  cobra.MaximumNArgs(1),
		RunE:  status.RunE(&statusCmdOpts),
	}
)

func init() {
	statusCmd.Flags().String("remote-storage-uri", "", "remote git repository URI")
	viper.BindPFlag(viperGitRemoteURI, statusCmd.Flags().Lookup("remote-storage-uri"))

	statusCmd.Flags().String("remote-storage-relative-data-file-path", "", "relative path to data file in remote git repository")
	viper.BindPFlag(viperGitRemoteRelativeDataFilePath, statusCmd.Flags().Lookup("remote-storage-relative-data-file-path"))
	viper.SetDefault(viperGitRemoteRelativeDataFilePath, "repoview.json")

	statusCmd.PreRun = func(cmd *cobra.Command, args []string) {
		statusCmdOpts.StoragePath = viper.GetString("storage.path")
		statusCmdOpts.RemoteDataRepositoryURI = viper.GetString(viperGitRemoteURI)
		statusCmdOpts.RemoteDataFilePath = viper.GetString(viperGitRemoteRelativeDataFilePath)
	}

	rootCmd.AddCommand(statusCmd)
}

const (
	viperGitRemoteURI                  = "GitRemote.URI"
	viperGitRemoteRelativeDataFilePath = "GitRemote.RelativeDataFilePath"
)
