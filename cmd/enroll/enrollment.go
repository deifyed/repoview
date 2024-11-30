package enroll

import "fmt"

type storer interface {
	StoreRepositoryPath(path string) error
}

func enroll(store storer, targetPath string) error {
	err := store.StoreRepositoryPath(targetPath)
	if err != nil {
		return fmt.Errorf("storing repository path: %w", err)
	}

	return nil
}
