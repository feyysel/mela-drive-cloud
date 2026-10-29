-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    first_name varchar(50) NOT NULL,
    last_name varchar(50),
    email varchar(255) NOT NULL UNIQUE,
    password_hash text NOT NULL,
    role varchar(20) NOT NULL CHECK (role IN ('owner', 'admin', 'member')),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS users;