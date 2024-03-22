CREATE TABLE IF NOT EXISTS admin (
    id SERIAL PRIMARY KEY,
    name TEXT,
    surname TEXT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);