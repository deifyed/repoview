package cmd

import (
	"github.com/deifyed/repoview/cmd/push"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pushCmd represents the push command
var (
	pushCmdOpts = push.Options{
		Fs: &afero.Afero{Fs: afero.NewOsFs()},
	}
	pushCmd = &cobra.Command{
		Use:   "push",
		Short: "Pushes statuses for all enrolled repositories to the data repository",
		RunE:  push.RunE(&pushCmdOpts),
	}
)

func init() {
	pushCmd.Flags().String("remote-storage-uri", "", "remote git repository URI")
	viper.BindPFlag(viperGitRemoteURI, pushCmd.Flags().Lookup("remote-storage-uri"))

	pushCmd.Flags().String("remote-storage-relative-data-file-path", "", "relative path to data file in remote git repository")
	viper.BindPFlag(viperGitRemoteRelativeDataFilePath, pushCmd.Flags().Lookup("remote-storage-relative-data-file-path"))
	viper.SetDefault(viperGitRemoteRelativeDataFilePath, "repoview.json")

	pushCmd.PreRun = func(cmd *cobra.Command, args []string) {
		pushCmdOpts.StoragePath = viper.GetString("storage.path")
		pushCmdOpts.RemoteDataRepositoryURI = viper.GetString(viperGitRemoteURI)
		pushCmdOpts.RemoteDataFilePath = viper.GetString(viperGitRemoteRelativeDataFilePath)
	}

	rootCmd.AddCommand(pushCmd)
}
