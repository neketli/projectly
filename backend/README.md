# Projectly backend

## 🛠️ Tools

- [Echo](https://pkg.go.dev/github.com/labstack/echo/v4/) - web framework
- [Pgx](https://pkg.go.dev/github.com/jackc/pgx/v5/) - postgresql driver
- [Squirrel](https://pkg.go.dev/github.com/Masterminds/squirrel) - SQL builder
- [Minio-go](https://pkg.go.dev/github.com/minio/minio-go/v7) - minio s3 connection
- [Zerolog](https://pkg.go.dev/github.com/rs/zerolog) - logger
- [Goose](https://pkg.go.dev/github.com/pressly/goose/v3) - migration tool
- [cleanenv](https://pkg.go.dev/github.com/ilyakaznacheev/cleanenv) - config reader
- [golang-jwt](https://pkg.go.dev/github.com/golang-jwt/jwt/v5) - JWT tokens
- [gothic](https://pkg.go.dev/github.com/markbates/goth/gothic) - OAuth tool

## 📑 Scheme

Check out [docs](./docs/) for logic database [schema](./docs/ER.jpg) and [plantuml](./docs/plantuml/) to understand how modules are related

## 📁 Project Structure

```bash
projectly
├─ backend            # Golang backend
│  ├─ cmd             # Application entry points
│  │  └─ server       # Backend HTTP server
│  ├─ config          # Project configuration files
│  ├─ docs            # Project docs
│  │  └─ swagger.yaml # HTTP Swagger/OpenAPI
│  ├─ go.mod          # Go module declaration
│  ├─ Makefile        # Backend make commands
│  ├─ internal        # App core
│  │  ├─ app          # Main entry point
│  │  └─ domain       # App domains
│  │     ├─ task          # Domain example - task
│  │     │  ├─ delivery   # HTTP Handlers
│  │     │  ├─ entity     # Domain entities (structs)
│  │     │  ├─ repository # Data saving
│  │     │  └─ usecase    # Core business logic of module
│  │     ├─ board
│  │     ├─ media
│  │     ├─ project
│  │     ├─ status
│  │     ├─ team
│  │     └─ user
│  ├─ migrations        # Database migrations
│  └─ pkg               # Reusable packages
```