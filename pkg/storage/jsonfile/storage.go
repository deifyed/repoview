package jsonfile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (s *Storage) load() error {
	if err := s.ensureStorageExists(); err != nil {
		return fmt.Errorf("ensuring storage exists: %w", err)
	}

	raw, err := s.Fs.ReadFile(s.absoluteStoragePath)
	if err != nil {
		return fmt.Errorf("reading records: %w", err)
	}

	if err := json.Unmarshal(raw, &s.records); err != nil {
		return fmt.Errorf("unmarshalling records: %w", err)
	}

	return nil
}

func (s *Storage) commit() error {
	if err := s.ensureStorageExists(); err != nil {
		return fmt.Errorf("ensuring storage exists: %w", err)
	}

	if s.records == nil {
		s.records = make(map[string]Record)
	}

	raw, err := json.Marshal(s.records)
	if err != nil {
		return fmt.Errorf("marshalling records: %w", err)
	}

	if err := s.Fs.WriteFile(s.absoluteStoragePath, raw, 0644); err != nil {
		return fmt.Errorf("writing records: %w", err)
	}

	return nil
}

func (s *Storage) ensureStorageExists() error {
	if err := s.ensureAbsoluteStoragePath(); err != nil {
		return fmt.Errorf("ensuring absolute storage path: %w", err)
	}

	if _, err := s.Fs.Stat(s.absoluteStoragePath); os.IsNotExist(err) {
		if err := s.Fs.MkdirAll(filepath.Dir(s.absoluteStoragePath), 0755); err != nil {
			return fmt.Errorf("creating storage directory: %w", err)
		}

		if err := s.commit(); err != nil {
			return fmt.Errorf("initializing storage: %w", err)
		}
	}
	return nil
}

func (s *Storage) ensureAbsoluteStoragePath() error {
	absPath, err := filepath.Abs(s.StoragePath)
	if err != nil {
		return fmt.Errorf("getting absolute path: %w", err)
	}

	s.absoluteStoragePath = absPath

	return nil
}
