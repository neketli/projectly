# Projectly

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://golang.org/)
[![Vue.js](https://img.shields.io/badge/Vue.js-3.5+-4FC08D?logo=vue.js)](https://vuejs.org/)

![preview](https://github.com/user-attachments/assets/c19fb3a0-86e8-449f-9a13-190676e22817)

**Projectly** is a modern open-source project management system designed to streamline task tracking and team collaboration. Ease of use and effective UI/UX make projectly a good tool for any team!

## ✨ Features

### 👥 Teams & Projects

- Collaborative team workspaces
- Project statistics and insights
- Kanban boards for task progress tracking
- Flexible project management tools

### 🔐 User Management

- Add users with customizable roles
- Editable user profiles
- Role-based access control

### ✅ Task Management

- Create and update task statuses
- Attach documents and files
- Commenting and mention system for communication
- Detailed progress tracking

## 🚀 One-command deployment

Deploy Projectly on any server with Docker:

### Official installer (recommended)

```bash
curl -fsSL https://projectly.ru/install.sh | sh
```

With a custom domain and HTTPS (Let's Encrypt):

```bash
curl -fsSL https://projectly.ru/install.sh | sh -s -- --domain app.example.com --email admin@example.com
```

Non-interactive / CI:

```bash
curl -fsSL https://projectly.ru/install.sh | sh -s -- --non-interactive --dir /opt/projectly --build-frontend
```

The installer will:

- Check the OS, install Docker via `get.docker.com` if missing, and verify Docker Compose v2
- Generate secure passwords and JWT secrets (`openssl rand`) into `.env` / `.env.backend` (mode `600`)
- Write `docker-compose.yml` and `nginx.conf`, optionally configuring HTTPS via Let's Encrypt (certbot)
- Pull images, run database migrations, create the MinIO bucket, and wait for health checks
- Print the app URL, MinIO console, credentials path, and management commands

| Flag | Description |
| --- | --- |
| `--dir <path>` | Install directory (default: `/opt/projectly` as root, `~/projectly` otherwise) |
| `--version <tag>` | Deploy a specific version/tag (default: `latest`) |
| `--domain <host>` | Public domain — enables HTTPS (Let's Encrypt) with `--email` |
| `--email <addr>` | Admin e-mail for Let's Encrypt registration/renewals |
| `--port <port>` | Host port for the web UI (default: `4242`) |
| `--build-frontend` | Build the frontend locally with your domain/port baked in (recommended for self-hosted installs) |
| `--no-tls` | Force plain HTTP even when `--domain` is given |
| `--non-interactive` | Never prompt; missing values fall back to defaults (CI) |
| `--force` | Overwrite existing credentials/configs without asking |
| `--yes` / `-y` | Assume "yes" for all confirmations |
| `--show-secrets` | Print generated credentials at the end |
| `--verify <sha256>` | Verify the script checksum before running |
| `--skip-docker-install` | Do not auto-install Docker |
| `--update` | Update an existing install (fetch configs, pull, restart). Domain and HTTPS settings are restored from the install marker unless overridden |
| `--uninstall` | Stop and remove containers (`--purge` also deletes data volumes) |
| `--dry-run` | Print what would be done without executing |

> **Integrity:** `docker-compose.yml` is verified against a published `sha256` checksum when fetched from GitHub; `nginx.conf` is always regenerated from the installer's local template. Run with `--verify <sha256>` to also check the script itself.

> **Note:** If the GHCR images require authentication, run `docker login ghcr.io -u YOUR_GITHUB_USERNAME` with a [PAT](https://github.com/settings/tokens) that has `read:packages` scope first.

### Requirements

- Linux (Ubuntu/Debian/CentOS/RHEL) or macOS (dev) with `curl`
- [Docker](https://docs.docker.com/engine/install/) + [Docker Compose](https://docs.docker.com/compose/install/) (v2) — installed automatically by the installer if missing
- `openssl` (pre-installed on macOS / Linux)

### Customization

Edit `.env` and `.env.backend` after the first run to change settings like domain, Yandex OAuth, or database credentials.

---

## 💡 Development

Follow these steps to set up a development environment.

## 🛠️ Tech Stack

### Frontend

- **Vue.js 3** - Progressive JavaScript framework
- **Nuxt 3** - For static site generation
- **Composition API** - As a modern approach to code splitting
- **TypeScript** - For types
- **Tailwind CSS** - Easy utility-first CSS styles
- **Element Plus** - Pretty UI component library

### Backend

- **Golang** - High-performance and easy to use
- **PostgreSQL** - Reliable relational database
- **MinIO** - Object storage
- **Docker** - Application containerization
- **Nginx** - Web server and reverse proxy

## 🚀 Quick start and development

Thank you for considering contributing to Projectly! Here are some steps to get you started:

### Prerequisites

Ensure you have the following installed:

- [Go 1.25+](https://golang.org/dl/)
- [Node.js 20+](https://nodejs.org/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Makefile](https://www.gnu.org/software/make/manual/make.html)

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/neketli/projectly.git
   cd projectly
   ```

2. **Configure Environment**

   Create a `.env` file in the project root and set the environment variables:

   ```env
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=[change_password]
    POSTGRES_DB=projectly

    MINIO_ROOT_USER=minioadmin
    MINIO_ROOT_PASSWORD=[change_password]
   ```

   Also you need to set `.env` in `./frontend`

   ```env
    NUXT_PUBLIC_API_HOST=http://localhost:8080 # or your API URL
    NUXT_PUBLIC_S3_HOST=http://localhost:8080
    NUXT_PUBLIC_SITE_URL=projectly.ru
    NUXT_APP_BASE_URL=/
   ```

   And do this again for `./backend`

   ```env
    APP_MODE=dev

    PG_DSN=postgresql://[postgres_user]:[postgres_password]@[host]:5432/projectly?sslmode=disable

    S3_HOST=[minio_host]
    S3_ACCESS=[minio_access_token]
    S3_SECRET=[minio_secret_token]

    AUTH_ACCESS_SECRET=[jwt_secret_encoding_key]
    AUTH_REFRESH_SECRET=[jwt_secret_encoding_key]
    
    AUTH_YANDEX_CLIENT_ID=
    AUTH_YANDEX_CLIENT_SECRET=
    AUTH_YANDEX_CALLBACK_URL=http://localhost:8080/api/v1/auth/yandex/callback

    SESSION_SECRET=[any_token_for_setup_sessions]
   ```

   [Setup yandex login](https://yandex.ru/dev/id/doc/ru/register-client)

3. **Set Up Infrastructure**

    ```bash
    docker compose up postgres minio nginx -d
    ```

4. **Install Dependencies**

   ```bash
   cd ./backend
   go mod download
   ```

   ```bash
    cd ../frontend
    npm install
   ```

5. **Make migrations**

    Make sure you have a `.env` in `./backend` and that you have created a bucket in your S3 provider with the correct access policies

    ```bash
    cd ./backend
    make migrate-up
    ```

6. **Run the Application**

   Backend:

   ```bash
   cd backend
   make build
   make run
   ```

   Frontend (in a new terminal):

   ```bash
   cd frontend
   npm run dev
   ```

7. **Access the Application**

   Open your browser and navigate to: `http://localhost:3000`

### Backend commands

```bash
# Build the application
make build

# Run the application
make run

# Generates swagger by swag-go doc comments
make swagger

# Applies up migrations to DB
make migrate-up

# Applies down migrations to DB
make migrate-down
```

### Frontend commands

```bash
# Development mode
npm run dev

# Build for production
npm run build

# Generate static files
npm run generate

# Preview production build
npm run preview

# Lint code with fixing
npm run lint
```

## 📁 Project Structure

```bash
projectly
├─ docker-compose.yml # Docker compose project config
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
├─ frontend          # Vue 3 (Nuxt.js) Frontend
│  ├─ app.vue        # Entry point
│  ├─ assets         # App static assets
│  ├─ components     # Reusable domain / UI components
│  ├─ composables    # Reusable logic / APIs
│  ├─ error.vue      # Fallback error page
│  ├─ i18n           # Locales
│  ├─ layouts        # Page base layouts
│  ├─ middleware     # Frontend middlewares
│  ├─ nuxt.config.ts # Nuxt config
│  ├─ pages          # App pages
│  ├─ plugins        # App plugins
│  ├─ public         # Public files
│  ├─ store          # Pinia stores
└─ └─ types          # App TS types
```

## 📞 Support

For questions or suggestions:

- 🐛 [Create an issue](https://github.com/neketli/projectly/issues)
- 📧 Contact the [author](https://github.com/neketli)

---

⭐ If you find this project useful, please give it a star on GitHub!
