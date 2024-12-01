package status

import (
	"fmt"
	"io"
	"strings"
)

const indent = "    "

func printStatusesForRepository(out io.Writer, uri string, machines []repositoryStatus) error {
	fmt.Fprint(out, strings.Trim(uri, "\n"))

	if len(machines) == 0 {
		fmt.Fprint(out, ": no data found\n")

		return nil
	}

	for _, machine := range machines {
		fmt.Fprintf(out, "\n%s%s", indent, strings.Trim(machine.MachineURI, "\n"))

		if machine.Status == "" {
			fmt.Fprint(out, " ⭐\n")
		} else {
			indented := strings.ReplaceAll(machine.Status, "\n", "\n"+indent+indent)
			trimmed := strings.TrimSuffix(indented, "\n")

			fmt.Fprintf(out, "\n%s%s\n", indent+indent, trimmed)
		}

	}

	return nil
}
