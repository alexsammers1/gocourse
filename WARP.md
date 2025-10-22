# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

Go course/learning repository.

## Development Commands

### Building
```bash
go build ./...
```

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests in a specific package
go test ./path/to/package

# Run a specific test
go test -run TestName ./path/to/package

# Run tests with coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Linting and Formatting
```bash
# Format code
go fmt ./...

# Run Go vet
go vet ./...

# Run golangci-lint (if installed)
golangci-lint run
```

### Dependencies
```bash
# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

## Architecture

*This section will be populated as the codebase grows.*
