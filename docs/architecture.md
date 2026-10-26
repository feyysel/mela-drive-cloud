# Architecture

## 1. System Overview

Mela Drive is a web-based file management platform. The frontend communicates with the backend API, which handles authentication, authorization, business logic, metadata persistence, and file operations.

## 2. Application Structure

| Directory           | Responsibility          |
| ------------------- | ----------------------- |
| `apps/api`          | Go backend and REST API |
| `apps/web`          | Next.js frontend        |
| `docs`              | Project documentation   |
| `.github/workflows` | CI workflows            |

Each application manages its own dependencies and configuration.

## 3. Core Components

| Component       | Responsibility                                            |
| --------------- | --------------------------------------------------------- |
| Web application | User interface and client-side interactions               |
| Backend API     | Business logic, validation, and authorization             |
| PostgreSQL      | File metadata, folders, sharing, and storage usage        |
| Redis           | Rate limiting, caching, and temporary data where required |
| Cloudflare R2   | File content and object storage                           |
| Mela Auth       | User authentication                                       |

## 4. Request Flow

1. The user interacts with the web application.
2. The frontend sends requests to the backend API.
3. The API authenticates the request and checks permissions.
4. The API reads or updates metadata in PostgreSQL.
5. File operations use Cloudflare R2 through the storage module.
6. The API returns the result to the frontend.

Redis supports applicable operations but is not the source of truth for persistent application data.

## 5. File Storage

Cloudflare R2 stores uploaded file content, while PostgreSQL stores the corresponding metadata and object keys.

A dedicated storage interface separates file-management logic from the storage provider. All file access must be authorized before content is delivered or a temporary access URL is issued.

## 6. Security and Authorization

- Integrate authentication with Mela Auth.
- Enforce authorization in the backend for every protected resource.
- Verify ownership or sharing permissions before file access.
- Validate uploads and enforce file size and type restrictions.
- Keep secrets in environment configuration or an approved secret manager.
- Use HTTPS in deployed environments.

## 7. Data Consistency

PostgreSQL and Cloudflare R2 operate independently and cannot participate in a single shared transaction.

The backend must handle failed uploads, metadata updates, and deletions safely. Cleanup and recovery procedures should address orphaned objects and metadata that references missing files.
