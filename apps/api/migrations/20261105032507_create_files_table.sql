-- +goose Up
CREATE TABLE files (
    id          uuid PRIMARY KEY,
    owner_id    uuid NOT NULL REFERENCES users (id),
    folder_id   uuid REFERENCES folders (id),
    name        text NOT NULL,
    storage_key text NOT NULL UNIQUE,
    size_bytes  bigint NOT NULL,
    mime_type   text NOT NULL,
    deleted_at  timestamptz,
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX files_owner_folder_idx ON files (owner_id, folder_id);

CREATE INDEX files_trash_idx ON files (deleted_at) WHERE deleted_at IS NOT NULL;

CREATE UNIQUE INDEX files_unique_name_idx
    ON files (owner_id, folder_id, lower(name))
    WHERE deleted_at IS NULL AND folder_id IS NOT NULL;

CREATE UNIQUE INDEX files_unique_root_name_idx
    ON files (owner_id, lower(name))
    WHERE deleted_at IS NULL AND folder_id IS NULL;

-- +goose Down
DROP TABLE files;