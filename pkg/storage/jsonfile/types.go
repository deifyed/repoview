package jsonfile

import "github.com/spf13/afero"

type Storage struct {
	Fs                  *afero.Afero
	StoragePath         string
	absoluteStoragePath string
	records             map[string]Record
}

type Record struct {
	AbsolutePath string `json:"absolute_path"`
}
