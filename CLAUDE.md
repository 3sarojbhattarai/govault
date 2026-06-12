# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
make build        # Build binary (output: govault)
make test         # Run all tests with verbose output
make fmt          # Format code
make lint         # Run go vet
make clean        # Remove build artifacts
make deps         # Tidy go.mod/go.sum
make run-dev      # Run without building (go run main.go)
```

To run a single test:
```bash
go test -v ./internal/crypto/... -run TestFunctionName
```

## Architecture

GoVault is a CLI password manager that encrypts a JSON vault file with a master password. The data flow is: CLI command → load vault file → decrypt → operate on in-memory model → re-encrypt → write back.

**Storage format**: `~/.govault/vault.enc` = `[32-byte random salt][AES-256-GCM encrypted JSON]`. The salt is stored unencrypted at the front of the file because it is needed to derive the key before decryption.

**Layers:**
- `cmd/root.go` — Cobra root command; calls `RegisterCommands` to attach subcommands
- `internal/commands/` — One file per subcommand (`add`, `get`, `list`, `delete`, `generate`). All share `getVaultPath`, `loadVault`, and `saveVault` helpers defined within the package
- `internal/models/models.go` — `Entry` (single credential) and `Vault` (slice of entries + salt bytes) structs with Add/Get/Delete/Update methods
- `internal/crypto/` — `crypto.go` does AES-256-GCM encrypt/decrypt; `key.go` does PBKDF2-HMAC-SHA256 key derivation (600,000 iterations) and salt generation
- `internal/storage/` — `Storage` interface (Load/Save/Exists/Delete); `FileStorage` implements it with atomic writes (write to temp file, then rename)

**Key constraint**: The `Vault` struct embeds the salt (`Salt []byte`) because it must travel with the data to re-derive the key on subsequent loads. The salt is stripped from the struct before JSON marshaling and prepended raw to the file instead.

**Dependencies**: Cobra (CLI), `golang.org/x/crypto` (PBKDF2 + bcrypt utilities), `golang.org/x/term` (hidden password prompt). No database, no network.
