CREATE TABLE IF NOT EXISTS gitscan.repositories (
    id serial PRIMARY KEY,
    name text NOT NULL,
    link text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
