package status

import (
	"fmt"
	"os"

	"github.com/deifyed/repoview/pkg/git"
	remotedata "github.com/deifyed/repoview/pkg/remote/git"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type Options struct {
	Fs          *afero.Afero
	StoragePath string

	RemoteDataRepositoryURI string
	RemoteDataFilePath      string
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting current working directory: %w", err)
		}

		if len(args) == 1 {
			targetPath = args[0]
		}

		repositoryURI, err := git.GetRepositoryURI(targetPath)
		if err != nil {
			return fmt.Errorf("getting repository URI: %w", err)
		}

		remote := remotedata.Remote{
			Fs:                   opts.Fs,
			RepositoryURI:        opts.RemoteDataRepositoryURI,
			RelativeDataFilePath: opts.RemoteDataFilePath,
		}

		remoteStatuses, err := getAllStatusesForRepository(&remote, repositoryURI)
		if err != nil {
			return fmt.Errorf("getting repository statuses: %w", err)
		}

		err = printStatusesForRepository(cmd.OutOrStdout(), repositoryURI, remoteStatuses)
		if err != nil {
			return fmt.Errorf("printing statuses: %w", err)
		}

		return nil
	}
}
