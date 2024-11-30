package enroll

import (
	"fmt"
	"path/filepath"

	"github.com/deifyed/repoview/pkg/fs"
	"github.com/deifyed/repoview/pkg/storage/jsonfile"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type EnrollOptions struct {
	Fs          *afero.Afero
	StoragePath string
}

func RunE(opts *EnrollOptions) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath, err := filepath.Abs(args[0])
		if err != nil {
			return fmt.Errorf("getting absolute path: %w", err)
		}

		if !fs.IsGitRepository(targetPath) {
			return fmt.Errorf("path %s is not a git repository", targetPath)
		}

		storage := &jsonfile.Storage{
			Fs:          opts.Fs,
			StoragePath: opts.StoragePath,
		}

		err = enroll(storage, targetPath)
		if err != nil {
			return fmt.Errorf("enrolling: %w", err)
		}

		return nil
	}
}
