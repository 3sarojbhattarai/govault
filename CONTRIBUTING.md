# Contributing to GoVault 🚀

First off, thank you for considering contributing to GoVault! It's people like you that make the open-source community such an amazing place to learn, inspire, and create.

## Prerequisites

To contribute to GoVault, you should have the following installed:
- [Go 1.21 or higher](https://go.dev/doc/install)
- [Git](https://git-scm.com/)
- `make` (optional, for using the Makefile)

## Project Setup

1. **Fork the repository** on GitHub.
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/govault.git
   cd govault
   ```
3. **Initialize Go modules** and download dependencies:
   ```bash
   make deps
   # or
   go mod tidy
   ```

## Development Workflow

We use a standard GitHub flow for development.

1. **Create a feature branch**:
   ```bash
   git checkout -b feature/amazing-feature
   ```
2. **Make your changes**: Implement your feature or fix.
3. **Format and Lint**: Ensure your code follows standards:
   ```bash
   make fmt
   make lint
   ```
4. **Test your changes**: Run the test suite:
   ```bash
   make test
   ```
5. **Commit your changes**: Write a clear, concise commit message:
   ```bash
   git commit -m 'Add amazing feature'
   ```
6. **Push to your fork**:
   ```bash
   git push origin feature/amazing-feature
   ```
7. **Open a Pull Request**: Submit your changes for review.

## Project Structure

GoVault follows a modular structure:

- `cmd/` - CLI command setup and configuration (Cobra)
- `internal/models/` - Data structures and business logic
- `internal/crypto/` - Encryption (AES-256-GCM) and key derivation (PBKDF2)
- `internal/storage/` - Vault persistence layer
- `internal/commands/` - CLI command implementations

## Coding Standards

- **Error Handling**: All operations should handle errors gracefully and provide meaningful messages to the user.
- **Modularity**: Keep modules decoupled and maintain a clean separation of concerns.
- **Documentation**: Update code comments and documentation (like this file) if you change functionality.
- **Tests**: Add unit tests for new features or bug fixes.

## Makefile Commands

The project includes a `Makefile` to simplify common development tasks:

- `make build`: Build the application binary.
- `make run-dev`: Run the application directly using `go run`.
- `make test`: Run all tests.
- `make fmt`: Format Go code.
- `make lint`: Run `go vet`.
- `make clean`: Remove built binaries.
- `make help`: Show all available commands.

## Security First 🔐

GoVault is a security-focused tool. When contributing, please:
- Never log sensitive information (passwords, keys).
- Use cryptographically secure random number generators for secrets.
- Ensure file permissions are handled correctly (e.g., 0600 for the vault).

---

**Built with ❤️ using Go**
