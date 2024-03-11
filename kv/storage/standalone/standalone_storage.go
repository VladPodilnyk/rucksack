package standalone

import (
	"github.com/vladpodilnyk/rucksack/kv/config"
	"github.com/vladpodilnyk/rucksack/kv/storage"
)

// Standalone is an implementation of `Storage` for a single-node of a database instance. It does not
// communicate with other nodes and all data is stored locally.
type StandaloneStorage struct {
	// TBD:
}

func NewStandaloneStorage(config *config.Config) *StandaloneStorage {
	return nil
}

func (s *StandaloneStorage) Start() error {
	return nil
}

func (s *StandaloneStorage) Stop() error {
	return nil
}

func (s *StandaloneStorage) Reader() (storage.StorageReader, error) {
	return nil, nil
}

func (s *StandaloneStorage) Write() error {
	return nil
}
