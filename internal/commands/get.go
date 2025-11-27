package commands

import (
	"fmt"
	"os"
	"syscall"

	"github.com/3sarojbhattarai/govault/internal/storage"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Retrieve a password entry",
	Long:  `Retrieve and display a password entry from the vault by name.`,
	Args:  cobra.ExactArgs(1),
	Run:   runGet,
}

func runGet(cmd *cobra.Command, args []string) {
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
		fmt.Fprintf(os.Stderr, "Error: Vault does not exist. Use 'add' to create your first entry.\n")
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

	// Get entry
	entry := vault.GetEntry(name)
	if entry == nil {
		fmt.Fprintf(os.Stderr, "Error: Entry '%s' not found\n", name)
		os.Exit(1)
	}

	// Display entry
	fmt.Println("\n═══════════════════════════════════════")
	fmt.Printf("  Name:       %s\n", entry.Name)
	fmt.Printf("  Username:   %s\n", entry.Username)
	fmt.Printf("  Password:   %s\n", entry.Password)
	if entry.URL != "" {
		fmt.Printf("  URL:        %s\n", entry.URL)
	}
	if entry.Notes != "" {
		fmt.Printf("  Notes:      %s\n", entry.Notes)
	}
	fmt.Printf("  Created:    %s\n", entry.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Modified:   %s\n", entry.ModifiedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("═══════════════════════════════════════\n")
}
