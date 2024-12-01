package git

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/afero"
)

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

func readData(fs *afero.Afero, targetFile string) (dataFile, error) {
	data := make(dataFile)

	raw, err := fs.ReadFile(targetFile)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling data: %w", err)
	}

	return data, nil
}

func writeData(fs *afero.Afero, targetFile string, data dataFile) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshalling data: %w", err)
	}

	err = fs.WriteFile(targetFile, raw, 0644)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}
