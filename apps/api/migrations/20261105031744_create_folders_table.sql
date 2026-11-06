-- +goose Up
CREATE TABLE folders (
    id CHAR(36) NOT NULL PRIMARY KEY,
    owner_id CHAR(36) NOT NULL,
    parent_id CHAR(36) NULL,
    name VARCHAR(255) NOT NULL,
    deleted_at DATETIME(6) NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    FOREIGN KEY (owner_id) REFERENCES users (id),
    FOREIGN KEY (parent_id) REFERENCES folders (id),
    INDEX folders_owner_parent_idx (owner_id, parent_id),
    INDEX folders_trash_idx (deleted_at),
    UNIQUE INDEX folders_unique_name_idx ((IF(deleted_at IS NULL, CONCAT(owner_id, ':', COALESCE(parent_id, 'root'), ':', name), NULL)))
);

-- +goose Down
DROP TABLE IF EXISTS folders;