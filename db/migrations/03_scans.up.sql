CREATE TABLE IF NOT EXISTS gitscan.scans (
    id serial PRIMARY KEY,
    repository_id int REFERENCES gitscan.repositories NOT NULL,
    status int NOT NULL DEFAULT 0,
    findings json,
    scanned_at TIMESTAMP, 
    finished_at TIMESTAMP, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
