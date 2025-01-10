# go-project

Run `go run cmd/main.go` to start the gin server

### Industry standard project structure
```
your-project/
├── cmd/                  # Entry points for the application
│   └── your-app/
│       └── main.go
├── config/               # Configuration files and logic
│   └── config.go
├── internal/             # Application-specific code
│   ├── controllers/      # HTTP handler logic
│   │   └── user_controller.go
│   ├── models/           # Data models and structs
│   │   └── user.go
│   ├── repositories/     # Data access logic
│   │   └── user_repository.go
│   ├── services/         # Business logic
│   │   └── user_service.go
│   ├── middlewares/      # Gin middlewares
│   │   └── auth_middleware.go
│   ├── db/               # Database connection and utilities
│   │   ├── db.go         # Database connection setup
│   │   └── migrations.go # Database migrations
│   └── integrations/     # Third-party API integrations
│       ├── api_client.go # Base API client
│       └── payment_api/  # Example: Payment API integration
│           ├── client.go # Payment client logic
│           └── types.go  # Request/Response structs
├── pkg/                  # Shared utility packages
│   ├── logger/           # Logging utilities
│   │   └── logger.go
│   ├── utils/            # General utilities (e.g., string, date helpers)
│   │   └── helpers.go
│   └── env/              # Environment variable helpers
│       └── env.go
├── routes/               # API route definitions
│   └── routes.go
├── migrations/           # Database migration files
├── .env                  # Environment variables
├── go.mod                # Go module file
├── go.sum                # Go dependency file
└── README.md             # Documentation
```
