package delist

import "fmt"

type storer interface {
	RemoveRepositoryPath(path string) error
}

func delist(store storer, targetPath string) error {
	err := store.RemoveRepositoryPath(targetPath)
	if err != nil {
		return fmt.Errorf("delisting repository path: %w", err)
	}

	return nil
}
