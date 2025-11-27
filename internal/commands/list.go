package commands

import (
	"fmt"
	"os"
	"syscall"

	"github.com/3sarojbhattarai/govault/internal/storage"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all password entries",
	Long:  `List all password entries stored in the vault (names only).`,
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
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

	// Display entries
	if len(vault.Entries) == 0 {
		fmt.Println("No entries found in vault.")
		return
	}

	fmt.Printf("\nFound %d entries:\n\n", len(vault.Entries))
	fmt.Println("═══════════════════════════════════════════════════════════════════")
	fmt.Printf("%-20s %-20s %-20s\n", "NAME", "USERNAME", "MODIFIED")
	fmt.Println("───────────────────────────────────────────────────────────────────")

	for _, entry := range vault.Entries {
		fmt.Printf("%-20s %-20s %-20s\n",
			truncate(entry.Name, 20),
			truncate(entry.Username, 20),
			entry.ModifiedAt.Format("2006-01-02 15:04"))
	}
	fmt.Println("═══════════════════════════════════════════════════════════════════\n")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
