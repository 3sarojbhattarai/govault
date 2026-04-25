# GoVault 🔐

A secure, modular password manager CLI built with Go.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Contributing](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

---

## Table of Contents

- [Features](#features)
- [Quick Start](#quick-start)
  - [Installation](#installation)
  - [Basic Usage](#basic-usage)
- [Commands](#commands)
  - [generate](#generate---generate-a-secure-password)
  - [add](#add---add-a-password-entry)
  - [list](#list---list-all-entries)
  - [get](#get---retrieve-a-password)
  - [delete](#delete---delete-an-entry)
- [Security](#security)
  - [How It Works](#how-it-works)
  - [Technical Details](#technical-details)
- [Contributing](#contributing)
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

# Build the application using Makefile
make build

# (Optional) Install globally
make install
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

---

### `list` - List All Entries

Display all stored password entries in a formatted table.

**Usage:**
```bash
./govault list
```

---

### `get` - Retrieve a Password

View the complete details of a password entry.

**Usage:**
```bash
./govault get <name>
```

---

### `delete` - Delete an Entry

Remove a password entry from the vault.

**Usage:**
```bash
./govault delete <name>
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

1. **Master Password**: You choose a master password that protects your entire vault.
2. **Key Derivation**: Your master password is converted to an encryption key using PBKDF2 with 600,000 iterations.
3. **Encryption**: All data is encrypted using AES-256-GCM before being saved.
4. **Storage**: Encrypted vault is stored at `~/.govault/vault.enc` with 0600 permissions.

### Technical Details

- **Algorithm**: AES-256-GCM (Galois/Counter Mode)
- **Key Derivation**: PBKDF2-HMAC-SHA256 (600,000 iterations)
- **Salt Size**: 256 bits (32 bytes)
- **Permissions**: 0600 (owner read/write only)

---

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines, project structure, and how to get started.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

