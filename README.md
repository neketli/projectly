# Projectly

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)](https://golang.org/)
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

## 💡 How to use

Install and prepare the project [by following this guide](#🚀-quick-start-and-development), then you can use [docker-compose](./docker-compose.yml) to easily set up projectly on your host.

Use latest version of projectly:

```bash
git clone https://github.com/neketli/projectly.git
docker compose up -d
```

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

- [Go 1.23+](https://golang.org/dl/)
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
    VITE_API_URL=http://localhost:8080 # or your API URL
    NUXT_PUBLIC_S3_HOST=https://projectly.ru
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
    
    # this will works only with domain (and not work locally btw)
    # for setup check links below
    AUTH_GOOGLE_CLIENT_ID=
    AUTH_GOOGLE_CLIENT_SECRET=
    AUTH_GOOGLE_CALLBACK_URL=http://localhost:8080/api/v1/auth/google/callback

    AUTH_YANDEX_CLIENT_ID=
    AUTH_YANDEX_CLIENT_SECRET=
    AUTH_YANDEX_CALLBACK_URL=http://localhost:8080/api/v1/auth/yandex/callback

    SESSION_SECRET=[any_token_for_setup_sessions]
   ```

    [Setup google login](https://developers.google.com/identity/sign-in/web/sign-in)

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
│  └─ types          # App TS types
└─ landing # Projectly.ru landing page
   ├─ app.vue        # Entry point
   ├─ assets         # Landing page static assets
   ├─ components     # Landing sections/UI components
   ├─ i18n           # Locales
   ├─ nuxt.config.ts # Nuxt config
   ├─ pages          # Landing pages
   └─ public         # Public files
```

## 📞 Support

For questions or suggestions:

- 🐛 [Create an issue](https://github.com/neketli/projectly/issues)
- 📧 Contact the [author](https://github.com/neketli)

---

⭐ If you find this project useful, please give it a star on GitHub!
