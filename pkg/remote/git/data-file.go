package git

import "fmt"

/*

{
  "github.com/deifyed/repoview": {
    statuses: {
      "localhost/user": "a\nb\nc"
    },
  }
}

*/

type repository struct {
	Statuses map[string]string `json:"statuses"`
}

type dataFile map[string]repository

func (d dataFile) upsertStatus(repositoryURI string, hostname string, username string, status string) error {
	d[repositoryURI].Statuses[fmt.Sprintf("%s/%s", hostname, username)] = status

	return nil
}
