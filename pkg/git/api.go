package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetRepositoryStatus(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "status", "--short")

	var stdout bytes.Buffer

	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("running git status: %w", err)
	}

	return stdout.String(), err
}

func GetRepositoryURI(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "remote", "get-url", "origin")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("running git remote get-url origin: %w", err)
	}

	uri := removeScheme(stdout.String())

	return strings.Trim(uri, "\n"), err
}
