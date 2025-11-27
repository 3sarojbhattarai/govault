# GoVault 🔐

A secure, modular password manager CLI built with Go.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Educational-blue.svg)](LICENSE)

---

## Table of Contents

- [Features](#features)
- [Quick Start](#quick-start)
  - [Installation](#installation)
  - [Basic Usage](#basic-usage)
- [Commands](#commands)
  - [generate - Generate Password](#generate---generate-a-secure-password)
  - [add - Add Entry](#add---add-a-password-entry)
  - [list - List Entries](#list---list-all-entries)
  - [get - Retrieve Entry](#get---retrieve-a-password)
  - [delete - Delete Entry](#delete---delete-an-entry)
- [Security](#security)
  - [How It Works](#how-it-works)
  - [Security Best Practices](#security-best-practices)
  - [Technical Details](#technical-details)
- [Development](#development)
  - [Prerequisites](#prerequisites)
  - [Project Setup](#project-setup)
  - [Development Mode](#development-mode)
  - [Production Build](#production-build)
  - [Installing Globally](#installing-globally)
- [Project Structure](#project-structure)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [Future Enhancements](#future-enhancements)
- [License](#license)

---

## Features

✨ **Secure Encryption**
- AES-256-GCM authenticated encryption
- PBKDF2 key derivation (600,000 iterations)
- Cryptographically secure random password generation

🏗️ **Modular Architecture**
- Clean separation of concerns
- Extensible storage interface
- Well-organized codebase

🛠️ **Full CLI Experience**
- Add, retrieve, list, and delete password entries
- Generate secure random passwords
- Interactive master password prompts
- Beautiful table-formatted output

---

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/3sarojbhattarai/govault
cd govault

# Download dependencies
go mod tidy

# Build the application
go build -o govault

# (Optional) Install globally
go install
```

### Basic Usage

```bash
# Generate a secure password
./govault generate

# Add a new password entry
./govault add github

# List all entries
./govault list

# Retrieve a password
./govault get github

# Delete an entry
./govault delete github
```

---

## Commands

### `generate` - Generate a Secure Password

Generate cryptographically secure random passwords.

**Usage:**
```bash
./govault generate [flags]
```

**Examples:**
```bash
# Default: 20 characters with all character types
./govault generate

# Custom length
./govault generate --length 32

# Exclude symbols
./govault generate --no-symbols

# Exclude digits
./govault generate --no-digits

# Exclude uppercase letters
./govault generate --no-uppercase

# Combine options
./govault generate -l 16 --no-symbols --no-digits
```

**Flags:**
- `-l, --length int` - Length of the password (default: 20)
- `--no-symbols` - Exclude symbols from password
- `--no-digits` - Exclude digits from password
- `--no-uppercase` - Exclude uppercase letters from password

---

### `add` - Add a Password Entry

Add a new password to your vault.

**Usage:**
```bash
./govault add <name>
```

**Interactive Prompts:**
- **Master password** - Creates vault on first use
- **Username** - Account username
- **Password** - Account password (hidden input)
- **URL** - Website URL (optional)
- **Notes** - Additional notes (optional)

**Example:**
```bash
./govault add github
# Enter master password: ********
# Username: user@example.com
# Password: ****************
# URL: https://github.com
# Notes: Personal account
# ✓ Entry 'github' added successfully
```

---

### `list` - List All Entries

Display all stored password entries in a formatted table.

**Usage:**
```bash
./govault list
```

**Output:**
```
Found 3 entries:

═══════════════════════════════════════════════════════════════════
NAME                 USERNAME             MODIFIED            
───────────────────────────────────────────────────────────────────
github               user@example.com     2025-11-27 18:20    
gmail                myemail@gmail.com    2025-11-27 18:21    
twitter              @myhandle            2025-11-27 18:22    
═══════════════════════════════════════════════════════════════════
```

---

### `get` - Retrieve a Password

View the complete details of a password entry.

**Usage:**
```bash
./govault get <name>
```

**Output:**
```
═══════════════════════════════════════
  Name:       github
  Username:   user@example.com
  Password:   MySecurePassword123!
  URL:        https://github.com
  Notes:      Personal account
  Created:    2025-11-27 18:20:15
  Modified:   2025-11-27 18:20:15
═══════════════════════════════════════
```

---

### `delete` - Delete an Entry

Remove a password entry from the vault.

**Usage:**
```bash
./govault delete <name>
```

**Example:**
```bash
./govault delete twitter
# Are you sure you want to delete 'twitter'? (yes/no): yes
# ✓ Entry 'twitter' deleted successfully
```

---

## Security

### How It Works

```
Master Password
      ↓
PBKDF2 (600,000 iterations + random salt)
      ↓
256-bit Encryption Key
      ↓
AES-256-GCM + unique nonce
      ↓
Encrypted Vault File (~/.govault/vault.enc)
```

1. **Master Password**: You choose a master password that protects your entire vault
2. **Key Derivation**: Your master password is converted to an encryption key using PBKDF2 with 600,000 iterations
3. **Encryption**: All data is encrypted using AES-256-GCM before being saved
4. **Storage**: Encrypted vault is stored at `~/.govault/vault.enc` with 0600 permissions

### Security Best Practices

- ✅ Choose a strong, unique master password
- ✅ Never share your master password
- ✅ Regularly backup your vault file (`~/.govault/vault.enc`)
- ✅ Store backups securely (encrypted external drive, etc.)
- ✅ Clear your terminal after viewing passwords
- ✅ Passwords are entered securely (not visible while typing)
- ⚠️ **Important**: If you forget your master password, your data **cannot be recovered**

### Technical Details

#### Encryption

- **Algorithm**: AES-256-GCM (Galois/Counter Mode)
- **Key Size**: 256 bits (32 bytes)
- **Authentication**: Built-in authenticated encryption (AEAD)
- **Nonce**: Unique per encryption operation

#### Key Derivation

- **Algorithm**: PBKDF2-HMAC-SHA256
- **Iterations**: 600,000
- **Salt Size**: 256 bits (32 bytes)
- **Output**: 256-bit encryption key

#### Storage

- **Location**: `~/.govault/vault.enc`
- **Format**: `[32 bytes salt][encrypted JSON data]`
- **Permissions**: 0600 (owner read/write only)
- **Atomic Writes**: Temporary file + rename for data safety

---

## Development

### Prerequisites

- Go 1.21 or higher
- Git (optional, for version control)

### Project Setup

1. **Clone or navigate to the project directory:**
   ```bash
   cd govault
   ```

2. **Initialize Go modules:**
   ```bash
   go mod tidy
   ```

   This downloads all required dependencies:
   - `github.com/spf13/cobra` - CLI framework
   - `golang.org/x/crypto` - Cryptography utilities
   - `golang.org/x/term` - Terminal password input

### Development Mode

#### Building the Application

```bash
# Standard build
go build -o govault

# Build with verbose output
go build -v -o govault
```

#### Running Without Building

```bash
# Run commands directly
go run main.go generate
go run main.go add mypassword
go run main.go list
go run main.go get mypassword
go run main.go delete mypassword
```

#### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

### Production Build

Build optimized binaries for production use:

```bash
# Build for current platform
go build -ldflags="-s -w" -o govault

# Build for Linux (from Mac/Windows)
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o govault-linux

# Build for Windows (from Mac/Linux)
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o govault.exe

# Build for macOS (from Linux/Windows)
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o govault-mac

# Build for macOS ARM (M1/M2)
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o govault-mac-arm64
```

> The `-ldflags="-s -w"` flags strip debug information and symbol tables to reduce binary size.

### Installing Globally

To install GoVault globally on your system:

```bash
# Install to $GOPATH/bin (make sure it's in your PATH)
go install

# Or copy the binary to a directory in your PATH
sudo cp govault /usr/local/bin/

# Make it executable (Linux/Mac)
sudo chmod +x /usr/local/bin/govault
```

After installation, you can run `govault` from anywhere.

---

## Project Structure

```
govault/
├── main.go                    # Application entry point
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── README.md                  # This file
├── LICENSE                    # License file
│
├── cmd/
│   └── root.go               # Cobra root command
│
└── internal/
    ├── models/
    │   └── models.go         # Data structures (Entry, Vault)
    │
    ├── crypto/
    │   ├── crypto.go         # AES-256-GCM encryption/decryption
    │   └── key.go            # PBKDF2 key derivation
    │
    ├── storage/
    │   ├── storage.go        # Storage interface
    │   └── file_storage.go   # File-based storage implementation
    │
    └── commands/
        ├── add.go            # Add password command
        ├── get.go            # Retrieve password command
        ├── list.go           # List passwords command
        ├── delete.go         # Delete password command
        └── generate.go       # Generate password command
```

### Module Organization

- **`cmd/`** - CLI command setup and configuration
- **`internal/models/`** - Data structures and business logic
- **`internal/crypto/`** - Encryption and key derivation
- **`internal/storage/`** - Vault persistence layer
- **`internal/commands/`** - CLI command implementations

---

## Troubleshooting

### "Incorrect master password" error

- Double-check your master password
- Make sure Caps Lock is off
- If you've forgotten it, there's **no recovery option**

### "Vault does not exist" error

- Use `govault add <name>` to create your first entry
- This will initialize the vault

### "Permission denied" errors

- Check that you have write permissions in your home directory
- Ensure `~/.govault/` directory has correct permissions (700)
- Verify vault file permissions: `ls -la ~/.govault/`

### "Corrupted vault data" error

- This may occur if:
  - The vault file was manually edited
  - The file transfer was interrupted
  - Different versions of GoVault were used
- Solution: Restore from backup or start fresh (delete `~/.govault/`)

---

## Contributing

Contributions are welcome! Please ensure:

1. **Follow existing code patterns** - Keep the modular structure
2. **Add error handling** - All operations should handle errors gracefully
3. **Test your changes** - Verify with various edge cases
4. **Update documentation** - Keep README and code comments current
5. **Keep modules decoupled** - Maintain clean separation of concerns

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Test thoroughly
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

---

## Future Enhancements

Potential features to consider:

- [ ] Password strength analyzer
- [ ] Vault export/import (encrypted)
- [ ] Multiple vault support
- [ ] Password history tracking
- [ ] Clipboard integration (auto-clear after timeout)
- [ ] Two-factor authentication
- [ ] Cloud sync (encrypted)
- [ ] Browser extensions
- [ ] Mobile apps (iOS/Android)
- [ ] Auto-fill integration
- [ ] Password sharing (encrypted)
- [ ] Audit log

---

## License

This project is for educational and personal use.

---

**Built with ❤️ using Go**

For questions or issues, please [open an issue](https://github.com/3sarojbhattarai/govault/issues) on GitHub.
