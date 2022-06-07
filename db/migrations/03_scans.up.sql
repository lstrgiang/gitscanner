CREATE TABLE IF NOT EXISTS gitscan.scans (
    id serial PRIMARY KEY,
    repository_id int REFERENCES gitscan.repositories NOT NULL;
    status int
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    scanned_at TIMESTAMP, 
    finished_at TIMESTAMP, 
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
