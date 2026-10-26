# Mela Drive

Mela Drive is a file management platform for securely storing, organizing, previewing, and sharing files.

## Project Structure

```text
.
├── apps/
│   ├── api/
│   └── web/
└── docs/

```

## Getting Started

### Prerequisites

- `Go`, `Node.js`, `pnpm`, `Just`, `PostgreSQL`, `Redis`.

Cloudflare R2 credentials and Mela Auth configuration are required for the corresponding features.

### Setup

**1. Clone the repository**

```bash
git clone https://github.com/feyysel/mela-drive-cloud.git
cd mela-drive-cloud
```

**2. Configure environment variables**

```bash
cp .env.example .env
```

**3. Install frontend dependencies**

```bash
cd apps/web
pnpm install
```

**4. Start PostgreSQL and Redis**

Ensure both services are running and configured in your environment.

**5. Run database migrations**

```bash
cd apps/api
just migrate
```

**6. Start the applications**

Run the API and web application using their respective development commands.

## Documentation

- [Specifications](docs/specifications.md)
- [Technology Stack](docs/tech-stack.md)
- [Architecture](docs/architecture.md)

## License

See [LICENSE](LICENSE).
