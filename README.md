# Go Workspace Multi-Module Project

A demonstration of Go workspace functionality with cross-module communication patterns using a shared common package.

## Overview

This project showcases Go's workspace feature (introduced in Go 1.18) with multiple modules that communicate through a shared package. It demonstrates how to structure a multi-module project where modules can depend on each other through a common bridge package.

## Architecture

```
┌─────────────────────────────────────────────┐
│              main.go (Root)                 │
│         Orchestrates both apps              │
└──────────────┬──────────────┬───────────────┘
               │              │
               ▼              ▼
        ┌──────────┐   ┌──────────┐
        │  app1    │   │  app2    │
        │  module  │   │  module  │
        └─────┬────┘   └──────────┘
              │             ▲
              │             │
              ▼             │
        ┌─────────────────┐ │
        │  common package │─┘
        │  (in root gwi)  │
        └─────────────────┘
```

### Dependency Flow

- **main.go** → calls both app1 and app2 directly
- **app1** → uses common package → calls app2
- **app2** → self-contained, no external dependencies
- **common** → bridges app1 to app2

## Project Structure

```
.
├── go.work              # Workspace configuration
├── go.mod               # Root module (gwi)
├── main.go              # Entry point
├── common/              # Shared bridge package
│   └── call.go          # Cross-module communication utilities
├── app1/                # First application module
│   ├── go.mod           # Module definition with gwi dependency
│   ├── main.go          # App1 entry point
│   └── app/
│       └── app.go       # App1 core functionality
└── app2/                # Second application module
    ├── go.mod           # Self-contained module definition
    ├── main.go          # App2 entry point
    └── app/
        └── app.go       # App2 core functionality
```

## Prerequisites

- **Go 1.25.1 or later** (project uses Go 1.25.1)
- Basic understanding of Go modules and workspaces

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```

2. Ensure Go is installed and accessible:
   ```bash
   go version
   # Should show: go version go1.25.x linux/amd64
   ```

3. The workspace is already configured with `go.work`. Verify it:
   ```bash
   cat go.work
   ```

## Usage

### Running the Main Application

```bash
go run main.go
```

**Expected Output:**
```
I'am APP1
I'am APP2
App2 said: I'am APP2
```

### Running Individual Apps

**App1:**
```bash
cd app1
go run main.go
```

**App2:**
```bash
cd app2
go run main.go
```

### Building Binaries

**Build all:**
```bash
# From root directory
go build -o bin/main ./main.go
go build -o bin/app1 ./app1/main.go
go build -o bin/app2 ./app2/main.go
```

**Run built binary:**
```bash
./bin/main
```

## Module Descriptions

### Root Module (`gwi`)

- **Purpose:** Contains the common package and main entry point
- **Package:** `common` - provides cross-module communication utilities
- **Location:** Root directory
- **Dependencies:** None (provides functionality to other modules)

### App1 Module (`gwi/app1`)

- **Purpose:** Demonstrates module that depends on root module
- **Key Function:** `WhoAmI()` - returns app1 identity
- **Key Function:** `WhoIsApp2()` - calls app2 through common package
- **Dependencies:** Root `gwi` module (via replace directive)
- **Pattern:** Uses shared common package to communicate with app2

### App2 Module (`gwi/app2`)

- **Purpose:** Self-contained module with no external dependencies
- **Key Function:** `WhoAmI()` - returns app2 identity
- **Dependencies:** None
- **Pattern:** Completely independent, called by others through common package

### Common Package (`gwi/common`)

- **Purpose:** Bridge package enabling app1 → app2 communication
- **Key Function:** `WhoIsApp2()` - wrapper that calls app2's WhoAmI function
- **Pattern:** Dependency injection bridge pattern

## How It Works

### Cross-Module Communication

The project demonstrates a sophisticated communication pattern:

1. **Direct Calls:** `main.go` can call both app1 and app2 directly
   ```go
   app1.WhoAmI()  // Direct call to app1
   app2.WhoAmI()  // Direct call to app2
   ```

2. **Indirect Calls via Common:** app1 calls app2 through the common package
   ```go
   app1.WhoIsApp2()  // app1 → common → app2
   ```

### Workspace Benefits

The `go.work` file enables:
- **Local Development:** Work on multiple modules simultaneously
- **No Publishing Required:** Test inter-module changes without publishing
- **Unified Dependencies:** Consistent dependency resolution across modules
- **Replace Directives:** app1's `go.mod` uses replace directive to reference local gwi module

### Module Replace Mechanism

In `app1/go.mod`:
```go
replace gwi => ../
require gwi v0.0.0-00010101000000-000000000000
```

This tells Go to use the local `../` directory instead of fetching from a remote source.

## Development

### Adding New Modules

1. Create new module directory:
   ```bash
   mkdir app3
   cd app3
   go mod init gwi/app3
   ```

2. Add to workspace:
   ```bash
   # From root directory
   go work use ./app3
   ```

3. Add dependencies as needed in the module's `go.mod`

### Modifying Common Package

When updating the common package:
1. Make changes in `common/` directory
2. Changes are immediately available to all modules in workspace
3. No need to publish or update versions during development

### Best Practices

- Keep `common` package focused on cross-cutting concerns
- Minimize dependencies in `common` to avoid coupling
- Use replace directives for local development
- Document any cross-module contracts in code comments
- Consider module boundaries carefully before adding dependencies

## Understanding the Code

### Main Entry Point (`main.go`)

```go
func main() {
    fmt.Println(app1.WhoAmI())      // Prints: I'am APP1
    fmt.Println(app2.WhoAmI())      // Prints: I'am APP2
    fmt.Println(app1.WhoIsApp2())   // Prints: App2 said: I'am APP2
}
```

The third call demonstrates the bridge pattern: app1 → common → app2.

### Common Bridge (`common/call.go`)

```go
func WhoIsApp2() string {
    return app2.WhoAmI()  // Calls into app2 module
}
```

This function acts as a bridge, allowing app1 to access app2 functionality without direct coupling.

### App1 Integration (`app1/app/app.go`)

```go
func WhoIsApp2() string {
    return "App2 said: " + common.WhoIsApp2()  // Uses bridge
}
```

App1 leverages the common package to communicate with app2, demonstrating loose coupling.

## Troubleshooting

### "command not found: go"

Ensure Go is in your PATH:
```bash
export PATH=$PATH:/usr/local/go/bin
```

### "cannot find module providing package"

Run from root directory:
```bash
go work sync
```

### Module version mismatch

Ensure all modules use compatible Go versions in their `go.mod` files.

## License

[Add your license here]

## Contributing

[Add contribution guidelines here]

## Contact

[Add contact information here]
