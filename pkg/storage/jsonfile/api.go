package jsonfile

import (
	"fmt"
)

func (s *Storage) StoreRepositoryPath(path string) error {
	if err := s.load(); err != nil {
		return fmt.Errorf("loading storage: %w", err)
	}

	s.records[path] = Record{AbsolutePath: path}

	if err := s.commit(); err != nil {
		return fmt.Errorf("committing storage: %w", err)
	}

	return nil
}

func (s *Storage) ListRepositoryPaths() ([]string, error) {
	if err := s.load(); err != nil {
		return nil, fmt.Errorf("loading storage: %w", err)
	}

	paths := make([]string, 0, len(s.records))

	for path := range s.records {
		paths = append(paths, path)
	}

	return paths, nil
}

func (s *Storage) RemoveRepositoryPath(path string) error {
	if err := s.load(); err != nil {
		return fmt.Errorf("loading storage: %w", err)
	}

	delete(s.records, path)

	if err := s.commit(); err != nil {
		return fmt.Errorf("committing storage: %w", err)
	}

	return nil
}
