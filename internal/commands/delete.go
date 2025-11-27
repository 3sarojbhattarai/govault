package commands

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/3sarojbhattarai/govault/internal/storage"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a password entry",
	Long:  `Delete a password entry from the vault by name.`,
	Args:  cobra.ExactArgs(1),
	Run:   runDelete,
}

func runDelete(cmd *cobra.Command, args []string) {
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

	if !store.Exists() {
		fmt.Fprintf(os.Stderr, "Error: Vault does not exist.\n")
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

	// Load vault
	vault, err := loadVault(store, string(masterPassword))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Check if entry exists
	if vault.GetEntry(name) == nil {
		fmt.Fprintf(os.Stderr, "Error: Entry '%s' not found\n", name)
		os.Exit(1)
	}

	// Confirm deletion
	fmt.Printf("Are you sure you want to delete '%s'? (yes/no): ", name)
	var confirmation string
	fmt.Scanln(&confirmation)

	if strings.ToLower(strings.TrimSpace(confirmation)) != "yes" {
		fmt.Println("Deletion cancelled.")
		return
	}

	// Delete entry
	if vault.DeleteEntry(name) {
		// Save vault
		if err := saveVault(store, vault, string(masterPassword)); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving vault: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Entry '%s' deleted successfully\n", name)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Failed to delete entry '%s'\n", name)
		os.Exit(1)
	}
}
