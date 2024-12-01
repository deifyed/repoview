package push

import (
	"fmt"

	"github.com/deifyed/repoview/pkg/remote/git"
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

		remote := &git.Remote{}

		err := push(storage, remote)
		if err != nil {
			return fmt.Errorf("pushing: %w", err)
		}

		return nil
	}
}
