# Valkey Sample Project

This project demonstrates a Go API service with Valkey integration, following the [Go Standard Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ja.md).

## Project Structure

```
.
├── cmd/api/core/            # Application entry point
│   └── main.go
├── internal/                # Private application code
│   ├── config/              # Configuration management
│   ├── handler/             # HTTP handlers
│   ├── router/              # Router configuration
│   └── feature/             # Feature modules
│       └── health/          # Health check utilities
├── docker/                  # Docker configuration
│   └── valkey/              # Valkey configuration
│       ├── local/           # Configuration files
│       │   └── valkey.conf
│       └── mnt/             # Persistent data
├── build/                   # Compiled binaries (generated)
├── go.mod                   # Go module file
├── Makefile                 # Build automation
└── README.md                # This file
```
