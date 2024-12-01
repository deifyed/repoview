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
