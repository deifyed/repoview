package git

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func readData(fs *afero.Afero, targetFile string) (dataFile, error) {
	data := make(dataFile)

	_, err := fs.Stat(targetFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return data, fmt.Errorf("statting file: %w", err)
		}

		return data, nil
	}

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
