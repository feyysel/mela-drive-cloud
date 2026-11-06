-- +goose Up
CREATE TABLE files (
    id CHAR(36) NOT NULL PRIMARY KEY,
    owner_id CHAR(36) NOT NULL,
    folder_id CHAR(36) NULL,
    name VARCHAR(255) NOT NULL,
    storage_key VARCHAR(512) NOT NULL UNIQUE,
    size_bytes BIGINT NOT NULL,
    mime_type VARCHAR(255) NOT NULL,
    deleted_at DATETIME(6) NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    FOREIGN KEY (owner_id) REFERENCES users (id),
    FOREIGN KEY (folder_id) REFERENCES folders (id),
    INDEX files_owner_folder_idx (owner_id, folder_id),
    INDEX files_trash_idx (deleted_at),
    UNIQUE INDEX files_unique_name_idx ((IF(deleted_at IS NULL, CONCAT(owner_id, ':', COALESCE(folder_id, 'root'), ':', name), NULL)))
);

-- +goose Down
DROP TABLE IF EXISTS files;