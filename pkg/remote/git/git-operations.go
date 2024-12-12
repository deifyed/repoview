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
	_, err := git(repositoryPath, "add", dataFilePath)
	if err != nil {
		return fmt.Errorf("adding: %w", err)
	}

	_, err = git(repositoryPath, "commit", "-m", message, dataFilePath)
	if err != nil {
		return fmt.Errorf("comitting: %w", err)
	}

	return nil
}

func pushChanges(repositoryPath string) error {
	_, err := git(repositoryPath, "push")
	if err != nil {
		return fmt.Errorf("pushing: %w", err)
	}

	return nil
}

func isRepositoryClean(repositoryPath string) (bool, error) {
	cmd := exec.Command("git", "-C", repositoryPath, "diff", "--quiet")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok && exiterr.ExitCode() == 1 {
			return false, nil
		}

		return false, fmt.Errorf("running git clone: %w: %s", err, stderr.String())
	}

	return true, nil
}

func git(repositoryPath string, args ...string) (string, error) {
	cmd := exec.Command("git", append([]string{"-C", repositoryPath}, args...)...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("running git: %w: %s %s", err, stderr.String(), stdout.String())
	}

	return stdout.String(), nil
}
