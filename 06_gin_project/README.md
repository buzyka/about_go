# Folder structure

```bash
monitoring/
├── build/                              # Compilation and build scripts
│   └── build_lambda.sh                 # Compile the Lambda function and zip it
│   └── run_amazon_linux.sh             # Run a Docker container with the `amazonlinux:2023` image.
│                                       # Inside this container, run `build_lambda.sh` to get a correctly
│                                       # compiled binary for AWS infrastructure.
├── cmd/
│   ├── web/                            # Entry point for the web application
│   │   └── main.go
│   ├── lambda/                         # Entry point for the AWS Lambda function
│   │   └── main.go
├── internal/
│   ├── domain/                         # Core domain layer
│   │   ├── entity/                     # Entities
│   │   ├── repository/                 # Repositories
│   │   ├── service/                    # Domain services
│   │   └── shared/                     # Shared domain concepts (e.g., errors, events) if needed
│   │       ├── errors.go
│   │       └── events.go
│   ├── application/                    # Application layer (use cases) if needed
│   ├── infrastructure/                 # Infrastructure layer (implementations)
|   │   ├── gocontainer/                # Container initialization (can be separated if necessary)
|   |   |   ├── gocontainer_web.go      # Web container initialization
│   │   │   └── gocontainer_lambda.go   # Lambda container initialization
|   │   ├── http/                       # Interfaces (web, Lambda handlers, etc.)
|   │   │   ├── web/                    # Web interface (HTTP handlers, routes)
|   │   │   │   ├── handlers/           # Handlers for HTTP requests
|   │   │   │   │   └── handler.go
|   │   │   │   ├── middleware.go       # HTTP middleware
|   │   │   │   └── routes.go           # Routes setup
|   │   │   ├── lambda/                 # AWS Lambda interface
|   │   │       └── handler.go          # Lambda handler for weather events
│   │   ├── repository/                 # Repository implementations, each repository in its own folder
│   │   │   └── weather_repository.go   # Weather repository implementation
│   │   ├── config/                     # Configuration management
│   │   │   └── config.go
│   │   ├── events/                     # Event processing (e.g., SNS, Kafka)
│   │   │   └── sns.go
│   │   └── logging/                    # Logging utilities
│   │       └── logger.go
├── pkg/                                # Optional: shared reusable packages
│   └── utils/                          # Common utilities
│       ├── http_helpers.go
│       └── validation.go
├── go.mod                              # Go module file
├── go.sum                              # Dependency checksum file
└── README.md                           # Documentation
```
