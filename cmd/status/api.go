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
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		remotedataRepositoryURI := "github.com/deifyed/repoview"

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
			Fs:            opts.Fs,
			RepositoryURI: remotedataRepositoryURI,
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

// TODO
// design:
// if command is run in a git repository, show the status of that repository in all hosts
// if command is run with --all-repositories
// if command is run with --all-local-repositories
/*
	storage := &jsonfile.Storage{
		Fs:          opts.Fs,
		StoragePath: opts.StoragePath,
	}

	paths, err := storage.ListRepositoryPaths()
	if err != nil {
		return fmt.Errorf("listing repository paths: %w", err)
	}

	for _, path := range paths {
		status, err := git.GetRepositoryStatus(path)
		if err != nil {
			return fmt.Errorf("getting repository status: %w", err)
		}

		uri, err := git.GetRepositoryURI(path)
		if err != nil {
			return fmt.Errorf("getting repository URI: %w", err)
		}

		err = printRepository(cmd.OutOrStdout(), uri, status)
		if err != nil {
			return fmt.Errorf("printing repository: %w", err)
		}
	}
*/
