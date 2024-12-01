package status

import (
	"github.com/deifyed/repoview/pkg/core"
)

type repositoryStatus struct {
	MachineURI string
	Status     string
}

type statusFetcher interface {
	GetRepositoryStatuses() ([]core.RepositoryStatus, error)
}

func getAllStatusesForRepository(remote statusFetcher, repositoryURI string) ([]repositoryStatus, error) {
	remoteStatuses, err := remote.GetRepositoryStatuses()
	if err != nil {
		return nil, err
	}

	relevantStatuses := make([]repositoryStatus, 0)

	for _, status := range remoteStatuses {
		if status.RepsitoryURI == repositoryURI {
			relevantStatuses = append(relevantStatuses, repositoryStatus{
				MachineURI: status.MachineURI,
				Status:     status.Status,
			})
		}
	}

	return relevantStatuses, nil
}
