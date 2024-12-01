package status

import (
	"fmt"

	"github.com/deifyed/repoview/pkg/git"
	"github.com/deifyed/repoview/pkg/storage/jsonfile"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type Options struct {
	Fs          *afero.Afero
	StoragePath string
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
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

		return nil
	}
}
