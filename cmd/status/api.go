package status

import (
	"github.com/deifyed/repoview/pkg/storage/jsonfile"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type StatusOptions struct {
	Fs          *afero.Afero
	StoragePath string
}

func RunE(opts *StatusOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		_ = &jsonfile.Storage{
			Fs:          opts.Fs,
			StoragePath: opts.StoragePath,
		}

		return nil
	}
}
