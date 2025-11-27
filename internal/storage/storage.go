package storage

// Storage defines the interface for vault persistence
type Storage interface {
	// Save saves encrypted vault data to storage
	Save(data []byte) error

	// Load loads encrypted vault data from storage
	Load() ([]byte, error)

	// Exists checks if the vault exists
	Exists() bool

	// Delete removes the vault from storage
	Delete() error
}
