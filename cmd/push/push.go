package push

import (
	"fmt"

	"github.com/deifyed/repoview/pkg/core"
	"github.com/deifyed/repoview/pkg/git"
)

type lister interface {
	ListRepositoryPaths() ([]string, error)
}

type uploader interface {
	UploadRepositoryStatus([]core.RepositoryStatus) error
}

func push(storage lister, remote uploader) error {
	repos, err := gatherStatuses(storage)
	if err != nil {
		return fmt.Errorf("gathering statuses: %w", err)
	}

	err = remote.UploadRepositoryStatus(repos)
	if err != nil {
		return fmt.Errorf("uploading statuses: %w", err)
	}

	return nil
}

func gatherStatuses(storage lister) ([]core.RepositoryStatus, error) {
	paths, err := storage.ListRepositoryPaths()
	if err != nil {
		return nil, err
	}

	statuses := make([]core.RepositoryStatus, len(paths))

	for index, path := range paths {
		status, err := git.GetRepositoryStatus(path)
		if err != nil {
			return nil, fmt.Errorf("getting repository status: %w", err)
		}

		uri, err := git.GetRepositoryURI(path)
		if err != nil {
			return nil, fmt.Errorf("getting repository uri: %w", err)
		}

		statuses[index] = core.RepositoryStatus{
			RepsitoryURI: uri,
			Status:       status,
		}
	}

	return statuses, nil
}
