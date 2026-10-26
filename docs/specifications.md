# Project Specifications

## 1. Overview

Mela Drive is an online file management platform for securely storing, organizing, previewing, and sharing files. It is part of the Mela ecosystem and integrates with Mela Auth for authentication.

### Objectives

- Provide a simple and reliable file management experience.
- Enable secure file sharing with configurable permissions.
- Maintain control over application data, storage, and infrastructure.
- Establish a reusable storage module for other Mela applications.
- Follow consistent development, security, and documentation standards.

## 2. Project Standards

### 2.1 Technology and Dependencies

- Use the project's agreed technology stack and supported stable releases.
- Use free, permissively licensed dependencies, preferably MIT, Apache-2.0, or BSD.
- Do not introduce paid APIs or trial-limited services without team approval.
- Maintain control over the database schema, migrations, and backups.
- Use PostgreSQL as the primary database and Redis for supporting workloads where appropriate.
- Keep file storage behind an independent interface to allow future storage-provider changes.

### 2.2 Collaboration

- Protect the `main` branch.
- Use feature branches for development.
- Submit changes through pull requests.
- Require at least one team-member review before merging.
- Use GitHub Issues and project boards to track work.
- Keep pull requests focused and include relevant tests and documentation updates.

### 2.3 Design and Accessibility

- Maintain consistent branding, typography, colors, and reusable components across Mela applications.
- Build responsive, mobile-first interfaces.
- Support English and Amharic from the beginning.
- Store interface translations in dedicated translation files.
- Provide accessible navigation, keyboard interaction, and form controls.
- Include appropriate loading, empty, success, and error states.

### 2.4 Documentation

Maintain the following documentation:

- Root `README.md` for the project overview and entry points.
- `docs/specifications.md` for functional and non-functional requirements.
- `docs/architecture.md` for architecture and design decisions.
- `docs/tech-stack.md` for technologies and dependency policies.
- `docs/development.md` for setup and development instructions.
- `docs/api.md` for API conventions and links to generated API documentation.
- `docs/user-guide.md` for end-user instructions.
- `.env.example` for documenting required environment variables without real secrets.

## 3. Functional Requirements

### 3.1 Authentication and Access Control

- Integrate with Mela Auth for user authentication.
- Identify the authenticated user securely on the server.
- Enforce authorization for every protected operation.
- Verify resource ownership or explicit sharing permissions before allowing access.
- Support view-only and edit permissions for shared resources.
- Reject unauthorized requests regardless of the controls displayed in the frontend.

### 3.2 File Management

Users must be able to:

- Upload single and multiple files.
- Upload files using drag-and-drop.
- Download files.
- Rename files.
- Move files between folders.
- Delete files by moving them to the trash.
- View file metadata, including name, size, type, and relevant timestamps.

The server must validate file types, file sizes, destination permissions, and available storage quota before accepting uploads.

### 3.3 Folder Management

Users must be able to:

- Create folders and nested folders.
- Navigate between folders using breadcrumb navigation.
- Move and rename folders.
- Delete folders.
- View the files and subfolders within a folder.

Folder operations must preserve consistent parent-child relationships and enforce authorization for all affected resources.

### 3.4 File Preview

Provide browser-based previews for supported formats:

- Images.
- PDF documents.
- Video.
- Audio.
- Plain text and source code.
- Office documents where a suitable free solution is available.

Preview access must follow the same authorization rules as file downloads. Office document previews must not expose private files to unauthorized third-party services.

### 3.5 Search and Organization

Users must be able to:

- Search files by name.
- Filter files by type and date.
- Sort files using relevant metadata.
- Switch between grid and list views.

Search results must include only resources the current user is authorized to access.

### 3.6 Trash and Recovery

- Move deleted files and folders to the trash.
- Allow users to restore deleted items.
- Allow permanent deletion from the trash.
- Permanently delete trashed items after 30 days.
- Apply retention consistently to folders and their contents.
- Ensure permanent deletion removes the associated stored objects and metadata according to the retention policy.

Cleanup operations must be safe to retry and must handle failures without silently leaving inconsistent data.

### 3.7 Storage Quotas

- Enforce a configurable storage quota for each user.
- Display used and available storage.
- Reject uploads that exceed the applicable quota.
- Handle concurrent uploads without allowing quota limits to be bypassed.
- Define how trashed files affect quota usage and apply that policy consistently.

The initial quota policy should be agreed upon before implementation, including whether trashed files continue to count toward usage.

### 3.8 File and Folder Sharing

Users must be able to:

- Generate shareable links for files and folders.
- Configure view-only or edit permissions.
- Set an optional expiration time.
- Revoke a share link when access is no longer required.

The server must validate each share link, its permissions, and its expiration. Revoked or expired links must no longer grant access.

Share tokens must be sufficiently unpredictable and must not expose internal database identifiers or private metadata.

### 3.9 Storage Module

File storage must be implemented as an independent module with a clearly defined interface.

The interface must support:

- Saving file content.
- Reading stored content.
- Deleting stored objects.
- Providing a secure mechanism for file previews.

The initial implementation may use local disk storage. The interface must allow a future self-hosted S3-compatible backend without requiring changes to file-management business logic.

Storage operations must use generated storage keys instead of trusting user-provided filenames.

## 4. Security Requirements

Security is mandatory across the application.

### 4.1 Authentication and Sessions

- Protect session credentials and tokens.
- Use HTTPS in deployed environments.
- Use secure, HttpOnly cookies when cookie-based sessions are implemented.
- Rate-limit sensitive authentication and public endpoints.

### 4.2 Input and Output Security

- Validate every request on the server.
- Use parameterized database queries.
- Encode output appropriately to prevent cross-site scripting.
- Configure an appropriate Content Security Policy.
- Enforce request size limits and timeouts.
- Avoid exposing stack traces, internal errors, or sensitive metadata to clients.

### 4.3 Authorization

- Enforce role-based access control where roles are applicable.
- Check resource ownership and sharing permissions on the server.
- Authorize downloads, previews, moves, renames, deletions, and sharing operations.
- Prevent unauthorized access through direct object identifiers or modified API requests.
- Ensure folder permissions are applied consistently to descendant resources according to the agreed sharing model.

### 4.4 File Upload Security

- Enforce file type and size restrictions on the server.
- Do not trust filenames, file extensions, or client-provided MIME types.
- Prevent path traversal and unauthorized file access.
- Keep stored files outside publicly accessible directories unless access is explicitly protected.
- Serve files only after checking authorization.
- Consider optional antivirus scanning with ClamAV.

### 4.5 Secrets and Dependencies

- Store secrets in environment variables or an approved secret-management system.
- Never commit credentials, private keys, or production secrets.
- Enable GitHub secret scanning and dependency vulnerability alerts where available.
- Review dependency licenses and security advisories.
- Keep dependencies updated and remove unnecessary packages.

### 4.6 Audit Logging

Record sensitive operations, including:

- Authentication events where available through the integration.
- Share-link creation and revocation.
- File and folder deletion.
- Permission changes.
- Other security-sensitive administrative actions.

Audit logs must not contain passwords, session credentials, share tokens, or file contents.

## 5. Non-Functional Requirements

### 5.1 Reliability

- Handle storage and database failures without reporting incomplete operations as successful.
- Make retryable operations safe to repeat where practical.
- Define cleanup and recovery behavior for interrupted uploads and deletions.
- Avoid assuming database transactions can atomically include filesystem operations.

### 5.2 Performance

- Use pagination for large file and folder listings.
- Avoid loading complete file contents into memory when streaming is appropriate.
- Apply sensible request limits and database query constraints.
- Keep file previews and large downloads from unnecessarily blocking unrelated requests.

### 5.3 Maintainability

- Keep API handlers, business logic, persistence, and storage concerns appropriately separated.
- Use versioned database migrations.
- Maintain consistent error handling and API response conventions.
- Add automated tests for important business rules and security boundaries.
- Document significant architectural decisions.

### 5.4 Localization

- Support English and Amharic.
- Keep user-facing text in translation resources.
- Avoid hardcoding interface strings in components.
- Ensure layouts accommodate translated text.

## 6. Optional Features

The following features are desirable but are not mandatory for the initial release:

- Sharing directly with Mela users by email.
- Resumable and chunked uploads for large files.
- File thumbnails.
- File version history.
- Starred files.
- Recently accessed files.
- ZIP downloads for folders.
- Optional antivirus scanning with ClamAV.

Prioritize these features through team planning rather than treating them as required for the first release.

## 7. Acceptance Criteria

A feature is considered complete when:

- Its agreed functional requirements are implemented.
- Server-side validation and authorization are in place.
- Relevant success, error, and edge cases are handled.
- Automated tests cover important behavior and security boundaries.
- Related documentation is updated.
- The change passes the repository's configured checks.
- The pull request has been reviewed and approved according to team policy.

The initial release is ready when the team has verified the agreed core workflows for authentication, file and folder management, previews, search, trash, quotas, sharing, and secure storage.
