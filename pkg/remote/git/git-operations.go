package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func clone(repositoryURI string, repositoryPath string) error {
	cmd := exec.Command("git", "clone", "--depth=1", repositoryURI, repositoryPath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("running git clone: %w: %s", err, stderr.String())
	}

	return nil
}

func commitFile(repositoryPath string, dataFilePath string, message string) error {
	commitCmd := exec.Command("git", "-C", repositoryPath, "commit", "-m", message, dataFilePath)

	err := commitCmd.Run()
	if err != nil {
		return fmt.Errorf("running git commit: %w", err)
	}

	return nil
}

func pushChanges(repositoryPath string) error {
	cmd := exec.Command("git", "-C", repositoryPath, "push")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("running git push: %w", err)
	}
	return nil
}
