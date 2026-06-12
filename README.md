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
- [Architecture](#architecture)
  - [Code Layers](#code-layers)
  - [Vault File Format](#vault-file-format)
  - [How the Master Password Works](#how-the-master-password-works)
- [Development](#development)
  - [Build & Run](#build--run)
  - [Commands](#commands)
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

**Install directly from GitHub (recommended):**

```bash
go install github.com/3sarojbhattarai/govault@latest
```

This downloads, builds, and places the binary in your `$GOPATH/bin`. Make sure `$GOPATH/bin` (usually `~/go/bin`) is on your `$PATH`.

**Or clone and build manually:**

```bash
git clone https://github.com/3sarojbhattarai/govault
cd govault
make build          # produces ./govault binary
make install        # installs to $GOPATH/bin
```

### Basic Usage

```bash
# Generate a secure password
govault generate

# Add a new password entry (interactive prompts follow)
govault add github

# List all stored entries
govault list

# Retrieve full details of an entry
govault get github

# Delete an entry
govault delete github
```

---

## Architecture

### Code Layers

```
┌──────────────────────────────────────────────────────┐
│                     CLI  (cmd/)                      │
│          Cobra root command + RegisterCommands       │
├──────────────────────────────────────────────────────┤
│              Commands  (internal/commands/)           │
│       add · get · list · delete · generate           │
│   (shared loadVault / saveVault helpers live here)   │
├─────────────────┬──────────────────┬─────────────────┤
│  Models         │  Crypto          │  Storage        │
│  (models/)      │  (crypto/)       │  (storage/)     │
│  Entry, Vault   │  AES-256-GCM     │  Storage iface  │
│  structs        │  PBKDF2 key deriv│  FileStorage    │
└─────────────────┴──────────────────┴─────────────────┘
```

Each command follows the same pattern: prompt for master password → `loadVault` (decrypt) → mutate in-memory `Vault` → `saveVault` (re-encrypt). The `Storage` interface decouples the commands from the filesystem, making alternative backends straightforward to add.

---

### Vault File Format

The vault is a single binary file at `~/.govault/vault.enc` (permissions `0600`):

```
┌───────────────────┬────────────────────────────────────────────┐
│  Salt  (32 bytes) │       AES-256-GCM encrypted payload        │
│   stored in plain │  [12-byte nonce │ ciphertext │ 16-byte tag]│
└───────────────────┴────────────────────────────────────────────┘
```

The salt must be readable before decryption (it is needed to re-derive the key), so it is stored unencrypted. The encrypted payload is the JSON-serialised `Vault` struct. The GCM authentication tag at the end of the payload lets the cipher detect any tampering or a wrong decryption key.

---

### How the Master Password Works

**The master password is never stored — not in the vault file, not on disk, not anywhere.**

**First use — vault creation:**

```
master password  ──┐
                   ├──▶  PBKDF2-HMAC-SHA256 (600,000 iter)  ──▶  256-bit key
random 32-byte     │
salt (generated)  ─┘
                                                                        │
                                             AES-256-GCM encrypt  ◀────┘
                                                        │
                                                        ▼
                              vault.enc = [ salt (32 B) | nonce | ciphertext | tag ]
```

**Every subsequent use:**

```
                              vault.enc = [ salt (32 B) | nonce | ciphertext | tag ]
                                                │
                                 read salt  ────┘

master password  ──┐
                   ├──▶  PBKDF2-HMAC-SHA256 (600,000 iter)  ──▶  256-bit key
stored salt       ─┘
                                                                        │
                                             AES-256-GCM decrypt  ◀────┘
                                                        │
                               wrong password?  GCM tag mismatch  ──▶  "incorrect master password"
                               correct password?  plaintext JSON   ──▶  Vault struct in memory
```

Because AES-GCM is authenticated encryption, a wrong password produces a tag-verification failure rather than silently decrypting garbage. This is how GoVault knows the password is wrong without ever storing a hash or hint.

The derived key and the password bytes exist only in process memory for the duration of the command and are released when the process exits.

---

## Development

### Build & Run

| Command | Description |
|---|---|
| `make build` | Compile binary (`./govault`) |
| `make run CMD="<args>"` | Run without building (e.g. `make run CMD="add github"`) |
| `make run-build CMD="<args>"` | Build then run (e.g. `make run-build CMD="get github"`) |
| `make test` | Run all tests with verbose output |
| `make fmt` | Format all Go source files |
| `make lint` | Run `go vet` |
| `make deps` | Tidy `go.mod` / `go.sum` |
| `make install` | Install binary to `$GOPATH/bin` |
| `make clean` | Remove all build artifacts |
| `make release` | Cross-compile for Linux, macOS (Intel + ARM), and Windows |

To run a single test:

```bash
go test -v ./internal/crypto/... -run TestFunctionName
```

---

### Commands

#### `generate` - Generate a Secure Password

Generate cryptographically secure random passwords.

```bash
govault generate [flags]
```

| Flag | Default | Description |
|---|---|---|
| `-l, --length` | `20` | Password length |
| `--no-symbols` | — | Exclude symbols |
| `--no-digits` | — | Exclude digits |
| `--no-uppercase` | — | Exclude uppercase letters |

---

#### `add` - Add a Password Entry

```bash
govault add <name>
```

Interactive prompts: master password · username · password (hidden) · URL (optional) · notes (optional). Creates the vault on first use.

---

#### `list` - List All Entries

```bash
govault list
```

Displays all entries in a formatted table.

---

#### `get` - Retrieve a Password

```bash
govault get <name>
```

Shows full details of a single entry.

---

#### `delete` - Delete an Entry

```bash
govault delete <name>
```

Prompts for confirmation before removing the entry.

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
