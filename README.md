# Portal Berita

**Project**: Portal Berita — a news portal built with **hexagonal architecture** (ports & adapters).

## Project Overview

Portal Berita is a news publishing platform that separates business logic from infrastructure concerns. The core domain contains pure business rules, while adapters handle external concerns like database access, HTTP request handling, and external services (Cloudflare). This design makes the codebase testable, replaceable, and maintainable.

## Architecture

### Hexagonal Architecture

- **Core domain** (`internal/core/domains`) — Entities and business rules, completely framework-agnostic.
- **Ports** (`internal/core`) — Interfaces that define contracts between the core and external layers. The core depends on these interfaces, not concrete implementations.
- **Adapters** (`internal/adapters`) — Concrete implementations of ports:
  - **Repositories** — Database access via GORM (PostgreSQL).
  - **Handlers** — HTTP request handlers that translate incoming requests to service calls.
  - **Cloudflare** — External service integration for CDN/security.
- **Services** (`internal/core/services`) — Orchestration layer that coordinates domain and port interactions.
- **Apps** (`internal/apps`) — Wiring and initialization of all components.
- **Config** (`configs/`) — Application configuration via Viper (JSON/YAML + environment variables).
- **Database** (`databases/`) — SQL migrations for PostgreSQL schema.

### Folder Structure

```
.
├── cmd/                    # CLI entry points (root, start)
│   ├── root.go            # Core API command
│   └── start.go           # Application startup
├── configs/               # Configuration
│   ├── config.go          # App & DB configuration
│   └── database.go        # PostgreSQL connection
├── databases/             # Database migrations
│   └── migrations/
│       ├── 000001_create_users_table.up.sql
│       ├── 000002_create_categories_table.up.sql
│       └── 000003_create_contents_table.up.sql
├── internal/              # Application code
│   ├── adapters/          # Port implementations
│   │   ├── cloudflare/    # Cloudflare integration
│   │   ├── handlers/      # HTTP handlers
│   │   └── repositories/  # Database repositories
│   ├── apps/              # App wiring
│   │   └── app.go         # Server startup
│   └── core/              # Business logic
│       ├── domains/       # Entities & rules
│       └── services/      # Service orchestration
├── libs/                  # Shared utilities
│   ├── conv/              # Conversion helpers
│   └── jwt/               # JWT handling
├── main.go                # Application entry point
└── go.mod / go.sum
```

## Key Components

- **Core Domain** — Pure business logic (entities, rules) with zero external dependencies.
- **Ports** — Interfaces that define how the core interacts with the outside world (e.g., `UserRepository`, `ContentRepository`, `CloudflareClient`).
- **Adapters** — Concrete implementations of ports:
  - `PostgresRepository` — GORM-based data access for users, categories, and contents.
  - `ContentHandler` — HTTP layer that receives requests and delegates to services.
  - `CloudflareAdapter` — External service integration.
- **Services** — Application-level coordination (e.g., `NewsService` that uses repositories and handlers).
- **Apps** — Composition root that wires everything together.
- **Config** — Centralized configuration via Viper (JSON/YAML + env vars).
- **Database** — PostgreSQL schema with migrations for users, categories, and news contents.

## Benefits

- **Testability** — Business logic can be unit tested without a database or HTTP server.
- **Flexibility** — Swap PostgreSQL for another DB, or replace Cloudflare with another CDN, without touching core.
- **Maintainability** — Clear separation of concerns makes code easier to understand and modify.

## Getting Started

```bash
# Run the application
go run main.go

# Or start via CLI
./cmd/start
```

## Database Migrations

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations.

### Install migrate CLI

Download the binary directly from GitHub releases:

```bash
# Linux (amd64)
wget -O migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz
tar -xzf migrate.tar.gz migrate
sudo mv migrate /usr/local/bin/migrate
rm migrate.tar.gz
```

```bash
# macOS (amd64)
wget -O migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.darwin-amd64.tar.gz
tar -xzf migrate.tar.gz migrate
sudo mv migrate /usr/local/bin/migrate
rm migrate.tar.gz
```

# Or download the precompiled binary directly from the releases page:
# https://github.com/golang-migrate/migrate/releases

### Run migrations

Create `~/.db/migrations` or use relative path:

```bash
# Up migrations (apply)
migrate -database postgres://$DATABASE_USER:$DATABASE_PASSWORD@localhost:$DATABASE_PORT/$DATABASE_NAME?sslmode=disable -path ./databases/migrations up

# Down migrations (rollback)
migrate -database postgres://$DATABASE_USER:$DATABASE_PASSWORD@localhost:$DATABASE_PORT/$DATABASE_NAME?sslmode=disable -path ./databases/migrations down

# Force version (set to specific state)
migrate -database postgres://$DATABASE_USER:$DATABASE_PASSWORD@localhost:$DATABASE_PORT/$DATABASE_NAME?sslmode=disable -path ./databases/migrations force 0
```

### Using with make or scripts

Add to your `Makefile`:

```makefile
migrate-up:
	migrate -database $(DATABASE_URL) -path ./databases/migrations up

migrate-down:
	migrate -database $(DATABASE_URL) -path ./databases/migrations down

migrate-force:
	migrate -database $(DATABASE_URL) -path ./databases/migrations force 0
```

## Configuration

Set environment variables or create a `.env` file with:
- `APP_PORT` — Port for the HTTP server.
- `APP_ENV` — Environment (`dev`, `prod`, etc.).
- `DATABASE_NAME`, `DATABASE_USER`, `DATABASE_HOST`, `DATABASE_PORT`, `DATABASE_PASSWORD`, `DATABASE_MAX_OPEN_POOL_CONNECTION`, `DATABASE_MAX_IDLE_CONNECTION`.
- `JWT_SECRET` — Secret key for JWT signing.
- `JWT_ISSUER` — Issuer identifier for JWTs.

## License

MIT
