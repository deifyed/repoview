package storage_test

import (
	"path/filepath"
	"testing"

	"github.com/deifyed/repoview/pkg/storage/jsonfile"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

type storager interface {
	StoreRepositoryPath(path string) error
	ListRepositoryPaths() ([]string, error)
}

type storageGenerator func(*afero.Afero, string) storager

func TestStoringRepositoryHappyPath(t *testing.T) {
	tests := []struct {
		name        string
		withStorage storageGenerator
		expected    string
		wantErr     bool
	}{
		{
			name: "happy path",
			withStorage: func(fs *afero.Afero, storagePath string) storager {
				return &jsonfile.Storage{
					Fs:          fs,
					StoragePath: storagePath,
				}
			},
			expected: "test",
			wantErr:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			storagePath := "/home/user/storage"
			repositoryPath := "/home/user/repos/repo_a"
			fs := &afero.Afero{Fs: afero.NewMemMapFs()}

			storage := tc.withStorage(fs, storagePath)

			err := fs.MkdirAll(filepath.Join(repositoryPath, ".git"), 0755)
			assert.NoError(t, err)

			err = storage.StoreRepositoryPath(repositoryPath)
			assert.NoError(t, err)

			repositories, err := storage.ListRepositoryPaths()
			assert.NoError(t, err)

			assert.Contains(t, repositories, repositoryPath)
		})
	}
}
