package commands

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultPasswordLength = 20
	uppercaseLetters      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters      = "abcdefghijklmnopqrstuvwxyz"
	digits                = "0123456789"
	symbols               = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

var (
	passwordLength int
	noSymbols      bool
	noDigits       bool
	noUppercase    bool
)

// GenerateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure random password",
	Long:  `Generate a cryptographically secure random password.`,
	Run:   runGenerate,
}

func init() {
	GenerateCmd.Flags().IntVarP(&passwordLength, "length", "l", defaultPasswordLength, "Length of the password")
	GenerateCmd.Flags().BoolVar(&noSymbols, "no-symbols", false, "Exclude symbols from password")
	GenerateCmd.Flags().BoolVar(&noDigits, "no-digits", false, "Exclude digits from password")
	GenerateCmd.Flags().BoolVar(&noUppercase, "no-uppercase", false, "Exclude uppercase letters from password")
}

func runGenerate(cmd *cobra.Command, args []string) {
	// Build character set
	charset := lowercaseLetters
	if !noUppercase {
		charset += uppercaseLetters
	}
	if !noDigits {
		charset += digits
	}
	if !noSymbols {
		charset += symbols
	}

	if len(charset) == 0 {
		fmt.Fprintf(os.Stderr, "Error: At least one character type must be enabled\n")
		os.Exit(1)
	}

	// Generate password
	password, err := generatePassword(passwordLength, charset)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nGenerated password: %s\n\n", password)
}

func generatePassword(length int, charset string) (string, error) {
	password := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}
