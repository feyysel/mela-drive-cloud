# Technology Stack

## Backend

| Technology    | Purpose                                  |
| ------------- | ---------------------------------------- |
| Go            | Programming language                     |
| Gin           | HTTP framework                           |
| PostgreSQL    | Primary database                         |
| `pgx/v5`      | PostgreSQL driver and connection pooling |
| Redis         | Caching and rate limiting                |
| `go-redis/v9` | Redis client                             |
| `redis_rate`  | Rate limiting                            |
| Goose         | Database migrations                      |

## Frontend

| Technology   | Purpose            |
| ------------ | ------------------ |
| Next.js      | Web framework      |
| React        | UI library         |
| TypeScript   | Static typing      |
| Tailwind CSS | Styling            |
| pnpm         | Package management |

## Development Tools

| Technology     | Purpose                              |
| -------------- | ------------------------------------ |
| Git            | Version control                      |
| GitHub         | Repository hosting and collaboration |
| Air            | Go hot reloading                     |
| GitHub Actions | CI workflows                         |

## Architecture and Infrastructure

| Area                | Approach                                                    |
| ------------------- | ----------------------------------------------------------- |
| Repository          | Plain monorepo with `apps/api` and `apps/web`               |
| Dependencies        | Managed independently by each application                   |
| Types               | Maintained separately in Go and TypeScript                  |
| File storage        | Cloudflare R2 (S3-compatible object storage)                |
| Persistent data     | PostgreSQL                                                  |
| Supporting services | Redis where required                                        |
| Hosting             | Managed hosting for application services and infrastructure |
