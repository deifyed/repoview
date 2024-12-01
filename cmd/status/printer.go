package status

import (
	"fmt"
	"io"
	"strings"
)

func printRepository(out io.Writer, uri string, status string) error {
	fmt.Fprint(out, strings.Trim(uri, "\n"))

	if status == "" {
		fmt.Fprint(out, " ⭐\n")

		return nil
	}

	indented := strings.ReplaceAll(status, "\n", "\n\t")
	trimmed := strings.TrimSuffix(indented, "\n")

	fmt.Fprintf(out, "\n\t%s\n", trimmed)

	return nil
}
