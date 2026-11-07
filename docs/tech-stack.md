# Technology Stack

## Backend

| Technology            | Purpose                                 |
| --------------------- | --------------------------------------- |
| Go                    | Programming language                    |
| Gin                   | HTTP framework                          |
| MySQL                 | Primary database                        |
| `go-sql-driver/mysql` | MySQL driver (used via `database/sql`)  |
| Redis                 | Caching and rate limiting               |
| `go-redis/v9`         | Redis client                            |
| `redis_rate`          | Rate limiting                           |
| Goose                 | Database migrations                     |
| Docker                | Containerized API builds and deployment |

## Frontend

| Technology   | Purpose            |
| ------------ | ------------------ |
| Next.js      | Web framework      |
| React        | UI library         |
| TypeScript   | Static typing      |
| Tailwind CSS | Styling            |
| pnpm         | Package management |

## Development Tools

| Technology     | Purpose                                       |
| -------------- | --------------------------------------------- |
| Git            | Version control                               |
| GitHub         | Repository hosting and collaboration          |
| Air            | Go hot reloading                              |
| Docker         | Local containers for the API and dependencies |
| GitHub Actions | CI workflows                                  |

## Architecture and Infrastructure

| Area                | Approach                                                    |
| ------------------- | ----------------------------------------------------------- |
| Repository          | Plain monorepo with `apps/api` and `apps/web`               |
| Dependencies        | Managed independently by each application                   |
| Types               | Maintained separately in Go and TypeScript                  |
| API packaging       | Docker image built from `apps/api`                          |
| File storage        | Cloudflare R2 (S3-compatible object storage)                |
| Persistent data     | MySQL                                                       |
| Supporting services | Redis where required                                        |
| Hosting             | Managed hosting for application services and infrastructure |
