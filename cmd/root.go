package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "govault",
	Short: "GoVault - A secure password manager CLI",
	Long: `GoVault is a secure, encrypted password manager for the command line.
	
It uses AES-256-GCM encryption to protect your passwords with a master password.
All data is stored locally in an encrypted vault file.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// RegisterCommands adds subcommands to the root command
func RegisterCommands(commands ...*cobra.Command) {
	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}
}
