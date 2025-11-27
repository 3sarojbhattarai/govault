package models

import "time"

// Entry represents a single password entry in the vault
type Entry struct {
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	URL        string    `json:"url,omitempty"`
	Notes      string    `json:"notes,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// Vault represents the encrypted vault containing all password entries
type Vault struct {
	Entries []Entry `json:"entries"`
	Salt    []byte  `json:"salt"`
}

// NewVault creates a new empty vault with a random salt
func NewVault(salt []byte) *Vault {
	return &Vault{
		Entries: make([]Entry, 0),
		Salt:    salt,
	}
}

// AddEntry adds a new entry to the vault
func (v *Vault) AddEntry(entry Entry) {
	v.Entries = append(v.Entries, entry)
}

// GetEntry retrieves an entry by name
func (v *Vault) GetEntry(name string) *Entry {
	for i := range v.Entries {
		if v.Entries[i].Name == name {
			return &v.Entries[i]
		}
	}
	return nil
}

// DeleteEntry removes an entry by name
func (v *Vault) DeleteEntry(name string) bool {
	for i, entry := range v.Entries {
		if entry.Name == name {
			v.Entries = append(v.Entries[:i], v.Entries[i+1:]...)
			return true
		}
	}
	return false
}

// UpdateEntry updates an existing entry
func (v *Vault) UpdateEntry(entry Entry) bool {
	for i := range v.Entries {
		if v.Entries[i].Name == entry.Name {
			v.Entries[i] = entry
			return true
		}
	}
	return false
}
