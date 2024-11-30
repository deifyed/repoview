package fs

import "os"

func IsGitRepository(path string) bool {
	gitPath := path + "/.git"
	if _, err := os.Stat(gitPath); os.IsNotExist(err) {
		return false
	}

	return true
}
