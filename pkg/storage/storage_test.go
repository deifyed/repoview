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
	}{
		{
			name: "jsonfile - happy path",
			withStorage: func(fs *afero.Afero, storagePath string) storager {
				return &jsonfile.Storage{
					Fs:          fs,
					StoragePath: storagePath,
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			storagePath := "/home/user/storage"

			withRepositories := []string{
				"/home/user/repos/repo_a", "/home/user/repos/repo_b", "/home/user/repos/repo_c",
			}

			fs := &afero.Afero{Fs: afero.NewMemMapFs()}
			storage := tc.withStorage(fs, storagePath)

			expectRespositories := make([]string, len(withRepositories))
			for _, repositoryPath := range withRepositories {
				err := fs.MkdirAll(filepath.Join(repositoryPath, ".git"), 0755)
				assert.NoError(t, err)

				expectRespositories = append(expectRespositories, repositoryPath)

				err = storage.StoreRepositoryPath(repositoryPath)
				assert.NoError(t, err)

				actualRepositories, err := storage.ListRepositoryPaths()
				assert.NoError(t, err)

				assert.Subset(t, withRepositories, actualRepositories)
			}
		})
	}
}
