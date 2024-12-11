# Projectly

![preview](https://github.com/user-attachments/assets/c19fb3a0-86e8-449f-9a13-190676e22817)

Projectly is an open source project management system. It helps teams to organize tasks, track progress, share files and collaborate on projects. Ease of use and effective UI/UX make projectly a good tool for any team!

## Features 🚀

- Teams and projects: collaborate with users in teams, get statistics by projects, use kanban boards to track tasks progress.

- User management: users can be added to the system and given different roles, each user has their own profile, which can be edited and managed.

- Task management: create tasks, update their status, attach documents and files, use mentions and comments to communicate with team.

## Tech stack ⚙️

- Frontend: Vue.js, composition API, typescript, tailwindcss, Element Plus
- Backend: Golang, PostgreSQL, Docker

## For contributors 🛠️

Thank you for considering contributing to Projectly! Here are some steps to get you started:

1. **Clone the repository**: Clone repository to your local machine using:

    ```bash
    git clone https://github.com/neketli/projectly.git
    ```

2. **Install deps**: You should install dependencies to run project

    ```bash
    cd <project_location>/projectly/backend && go mod download
    cd <project_location>/projectly/frontend && npm i
    ```

3. **Setup infra**: For development you will need infrastructure and setup configs (.env)

    ```bash
    cd <project_location>/projectly && docker compose up postgres minio nginx
    ```

4. **Run project**: Now you available to run project

    ```bash
    cd <project_location>/projectly/backend && make build && make run
    cd <project_location>/projectly/frontend && npm run dev
    ```
