-- +goose Up
CREATE TABLE folders (
    id         uuid PRIMARY KEY,
    owner_id   uuid NOT NULL REFERENCES users (id),
    parent_id  uuid REFERENCES folders (id),
    name       text NOT NULL,
    deleted_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX folders_owner_parent_idx ON folders (owner_id, parent_id);

CREATE INDEX folders_trash_idx ON folders (deleted_at) WHERE deleted_at IS NOT NULL;

CREATE UNIQUE INDEX folders_unique_name_idx
    ON folders (owner_id, parent_id, lower(name))
    WHERE deleted_at IS NULL AND parent_id IS NOT NULL;

CREATE UNIQUE INDEX folders_unique_root_name_idx
    ON folders (owner_id, lower(name))
    WHERE deleted_at IS NULL AND parent_id IS NULL;

-- +goose Down
DROP TABLE folders;