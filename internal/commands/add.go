package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/3sarojbhattarai/govault/internal/crypto"
	"github.com/3sarojbhattarai/govault/internal/models"
	"github.com/3sarojbhattarai/govault/internal/storage"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new password entry",
	Long:  `Add a new password entry to the vault. You will be prompted for the entry details.`,
	Args:  cobra.ExactArgs(1),
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	name := args[0]

	// Get vault path
	vaultPath, err := getVaultPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Initialize storage
	store, err := storage.NewFileStorage(vaultPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing storage: %v\n", err)
		os.Exit(1)
	}

	// Get master password
	fmt.Print("Enter master password: ")
	masterPassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading password: %v\n", err)
		os.Exit(1)
	}

	// Load or create vault
	var vault *models.Vault
	if store.Exists() {
		vault, err = loadVault(store, string(masterPassword))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading vault: %v\n", err)
			os.Exit(1)
		}

		// Check if entry already exists
		if vault.GetEntry(name) != nil {
			fmt.Fprintf(os.Stderr, "Error: Entry '%s' already exists\n", name)
			os.Exit(1)
		}
	} else {
		// Create new vault
		salt, err := crypto.GenerateSalt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating salt: %v\n", err)
			os.Exit(1)
		}
		vault = models.NewVault(salt)
	}

	// Prompt for entry details
	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading password: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("URL (optional): ")
	var url string
	fmt.Scanln(&url)

	fmt.Print("Notes (optional): ")
	var notes string
	fmt.Scanln(&notes)

	// Create entry
	entry := models.Entry{
		Name:       name,
		Username:   username,
		Password:   string(password),
		URL:        url,
		Notes:      notes,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	// Add to vault
	vault.AddEntry(entry)

	// Save vault
	if err := saveVault(store, vault, string(masterPassword)); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving vault: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Entry '%s' added successfully\n", name)
}

func getVaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".govault", "vault.enc"), nil
}

func loadVault(store storage.Storage, masterPassword string) (*models.Vault, error) {
	// Load encrypted data
	data, err := store.Load()
	if err != nil {
		return nil, err
	}

	// Format: [32 bytes salt][encrypted vault data]
	if len(data) < crypto.SaltSize {
		return nil, fmt.Errorf("corrupted vault data")
	}

	// Extract salt (first 32 bytes)
	salt := data[:crypto.SaltSize]
	encryptedData := data[crypto.SaltSize:]

	// Derive key from master password and salt
	key := crypto.DeriveKey(masterPassword, salt)

	// Decrypt the vault data
	decryptedData, err := crypto.Decrypt(encryptedData, key)
	if err != nil {
		return nil, fmt.Errorf("incorrect master password")
	}

	// Parse decrypted vault JSON
	var vault models.Vault
	if err := json.Unmarshal(decryptedData, &vault); err != nil {
		return nil, fmt.Errorf("corrupted vault data")
	}

	// Restore the salt
	vault.Salt = salt

	return &vault, nil
}

func saveVault(store storage.Storage, vault *models.Vault, masterPassword string) error {
	// Derive key
	key := crypto.DeriveKey(masterPassword, vault.Salt)

	// Marshal vault (entries only, not the salt)
	vaultJSON, err := json.Marshal(vault)
	if err != nil {
		return err
	}

	// Encrypt the vault data
	encryptedData, err := crypto.Encrypt(vaultJSON, key)
	if err != nil {
		return err
	}

	// Combine salt + encrypted data
	// Format: [32 bytes salt][encrypted vault data]
	data := append(vault.Salt, encryptedData...)

	// Save to storage
	return store.Save(data)
}
