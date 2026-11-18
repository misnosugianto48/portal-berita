# Portal Berita

**Project**: Portal Berita — a modern news portal API built with **Go** using **Hexagonal Architecture** (Ports & Adapters) and **Fiber v3**.

## Project Overview

Portal Berita is a news publishing platform that strictly separates business logic from infrastructure concerns. The core domain contains pure business rules and models, while adapters handle external concerns such as database access, HTTP request handling, authentication, and cloud storage (Cloudflare R2). This design keeps the codebase testable, modular, and maintainable.

---

## Tech Stack

- **Language**: Go 1.26
- **Web Framework**: [Fiber v3](https://github.com/gofiber/fiber/v3)
- **CLI & Configuration**: [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper)
- **Database & ORM**: PostgreSQL via [GORM](https://gorm.io) (PgBouncer transaction-pooler compatible)
- **Database Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Authentication**: JWT ([golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)) & Bcrypt
- **Validation**: [go-playground/validator/v10](https://github.com/go-playground/validator) integrated directly into Fiber's `StructValidator`
- **Cloud Storage**: AWS SDK Go v2 configured for Cloudflare R2
- **Logging**: [zerolog](https://github.com/rs/zerolog)

---

## Architecture & Folder Structure

```
.
├── cmd/                          # CLI commands (Cobra)
│   ├── root.go                   # Root command & Viper config loader
│   └── start.go                  # Server startup command
├── configs/                      # Infrastructure & runtime configs
│   ├── cloudflareR2.go           # AWS S3 / Cloudflare R2 client setup
│   ├── config.go                 # App, database, and cloud storage config struct
│   ├── database.go               # PostgreSQL connection with connection pooling
│   └── validator.go              # Fiber v3 StructValidator with user-friendly messages
├── databases/
│   ├── migrations/               # SQL migrations for golang-migrate
│   │   ├── 000001_create_users_table
│   │   ├── 000002_create_categories_table
│   │   └── 000003_create_contents_table
│   └── seeders/                  # Initial database seeds (Admin user)
├── internal/
│   ├── adapters/                 # Outgoing & incoming adapters (Ports implementation)
│   │   ├── cloudflare/           # Cloudflare R2 storage adapter
│   │   ├── handlers/             # HTTP controller handlers (Fiber v3)
│   │   │   ├── requests/         # DTO request structs & validation tags
│   │   │   ├── responses/        # API response envelopes
│   │   │   └── auth_handler.go   # Authentication handler
│   │   └── repositories/         # Database persistence with GORM
│   │       └── auth_repository.go
│   ├── apps/
│   │   └── app.go                # Composition root (wiring DI, middlewares, routes)
│   └── core/
│       ├── domains/              # Pure domain logic (entities & models)
│       │   ├── entity/           # Domain entities (JWT, Pagination, User)
│       │   └── models/           # GORM database models (User, Category, Content)
│       └── services/             # Core business orchestration & port interfaces
│           └── auth_service.go
├── libs/                         # Reusable shared libraries
│   ├── auth/                     # JWT token generator & verifier
│   ├── conv/                     # Password hashing, slug generator, type converters
│   ├── middlewares/              # Fiber middleware (Bearer token authentication)
│   └── paginations/              # Pagination metadata calculator
├── main.go                       # Application entrypoint
├── go.mod / go.sum
└── README.md
```

---

## Request Validation & Binding

This project leverages Fiber v3's `c.Bind().JSON(&req)` combined with a custom `StructValidator` in [configs/validator.go](configs/validator.go):

- **Automatic Validation**: Simply calling `c.Bind().JSON(&dto)` automatically unmarshals JSON and evaluates all `validate:"..."` struct tags.
- **User-Friendly Error Messages**: Field names use their JSON tag name (`email` instead of `Email`) and error messages are human-readable (e.g. `email is required, password must be at least 8 characters`) rather than raw validator internal dumps.

---

## Getting Started

### Prerequisites

- Go 1.22+ (tested on Go 1.26)
- PostgreSQL database
- (Optional) [golang-migrate CLI](https://github.com/golang-migrate/migrate)

### 1. Environment Configuration

Create a `.env` file in the root directory:

```env
# Application
APP_PORT=8080
APP_ENV=development

# JWT Authentication
JWT_SECRET_KEY=your-secret-key-here
JWT_ISSUER=portal-berita

# PostgreSQL Database
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_NAME=portal_berita
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_MAX_OPEN_POOL_CONNECTION=10
DATABASE_MAX_IDLE_CONNECTION=5

# Cloudflare R2 Storage (Optional)
CLOUDFLARE_R2_BUCKET_NAME=my-bucket
CLOUDFLARE_R2_API_KEY=your-r2-key
CLOUDFLARE_R2_API_SECRET=your-r2-secret
CLOUDFLARE_R2_ACCOUNT_ID=your-account-id
CLOUDFLARE_R2_PUBLIC_URL=https://pub-xxx.r2.dev
```

> **Note for Supabase / PgBouncer**: Database connection string automatically includes `default_query_exec_mode=exec` to support transaction poolers without prepared statement conflicts.

### 2. Database Migrations

Apply migrations using `golang-migrate`:

```bash
# Apply all up migrations
migrate -database "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -path ./databases/migrations up

# Rollback migrations
migrate -database "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -path ./databases/migrations down
```

### 3. Run the Application

```bash
# Run directly
go run main.go

# Or run via CLI command
go run main.go start
```

---

## API Endpoints

### Authentication
- `POST /api/v1/login` — Authenticate admin user and receive JWT access token.
  - Body:
    ```json
    {
      "email": "admin@email.com",
      "password": "admin123"
    }
    ```

---

## License

MIT
