package git

import (
	"strings"
)

func removeScheme(uri string) string {
	result := strings.Split(uri, "//")

	if len(result) == 1 {
		return result[0]
	}

	return result[1]
}
