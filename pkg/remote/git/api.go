package git

import (
	"fmt"

	"github.com/deifyed/repoview/pkg/core"
	"github.com/spf13/afero"
)

const (
	statusFileName = "repoview.json"
)

type Remote struct {
	Fs            *afero.Afero
	RepositoryURI string
}

func (r *Remote) UploadRepositoryStatus(statuses []core.RepositoryStatus) error {
	info, err := generateLocalInfo(r.Fs, r.RepositoryURI)
	if err != nil {
		return fmt.Errorf("generating local info: %w", err)
	}

	err = clone(r.RepositoryURI, info.RepositoryPath)
	if err != nil {
		return fmt.Errorf("cloning repository: %w", err)
	}

	existingData, err := readData(r.Fs, info.DataFilePath)
	if err != nil {
		return fmt.Errorf("loading data: %w", err)
	}

	for _, repo := range statuses {
		err = existingData.upsertStatus(repo.RepsitoryURI, info.Hostname, info.Username, repo.Status)
		if err != nil {
			return fmt.Errorf("updating data: %w", err)
		}
	}

	err = writeData(r.Fs, info.DataFilePath, existingData)
	if err != nil {
		return fmt.Errorf("writing data: %w", err)
	}

	err = commitFile(info.RepositoryPath, info.DataFilePath, "Updated status")
	if err != nil {
		return fmt.Errorf("adding and committing: %w", err)
	}

	err = pushChanges(info.RepositoryPath)
	if err != nil {
		return fmt.Errorf("pushing: %w", err)
	}

	return nil
}
