package git

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/afero"
)

type localInfo struct {
	Hostname       string
	Username       string
	RepositoryPath string
	DataFilePath   string
}

func generateLocalInfo(fs *afero.Afero, relativeDataFilePath string) (*localInfo, error) {
	localHostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("getting hostname: %w", err)
	}

	localUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("getting current user: %w", err)
	}

	repositoryPath, err := fs.TempDir("", "repoview")
	if err != nil {
		return nil, fmt.Errorf("creating temporary directory: %w", err)
	}

	dataFilePath := filepath.Join(repositoryPath, relativeDataFilePath)

	return &localInfo{
		Hostname:       localHostname,
		Username:       localUser.Username,
		RepositoryPath: repositoryPath,
		DataFilePath:   dataFilePath,
	}, nil
}
