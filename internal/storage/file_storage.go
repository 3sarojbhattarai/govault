package storage

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileStorage implements Storage interface using local file system
type FileStorage struct {
	path string
}

// NewFileStorage creates a new file-based storage
func NewFileStorage(path string) (*FileStorage, error) {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, err
	}

	return &FileStorage{
		path: path,
	}, nil
}

// Save saves encrypted vault data to file
func (fs *FileStorage) Save(data []byte) error {
	// Write to temporary file first for atomic operation
	tempPath := fs.path + ".tmp"
	if err := ioutil.WriteFile(tempPath, data, 0600); err != nil {
		return err
	}

	// Rename to actual file (atomic on most systems)
	if err := os.Rename(tempPath, fs.path); err != nil {
		os.Remove(tempPath) // Clean up temp file
		return err
	}

	return nil
}

// Load loads encrypted vault data from file
func (fs *FileStorage) Load() ([]byte, error) {
	if !fs.Exists() {
		return nil, errors.New("vault does not exist")
	}

	data, err := ioutil.ReadFile(fs.path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Exists checks if the vault file exists
func (fs *FileStorage) Exists() bool {
	_, err := os.Stat(fs.path)
	return err == nil
}

// Delete removes the vault file
func (fs *FileStorage) Delete() error {
	if !fs.Exists() {
		return errors.New("vault does not exist")
	}

	return os.Remove(fs.path)
}
