CREATE DATABASE url_shortener;
\c url_shortener;
CREATE TABLE url_infos (
    id SERIAL PRIMARY KEY,
    long_url TEXT,
    short_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    deleted_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL
);