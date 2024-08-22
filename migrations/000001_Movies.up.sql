CREATE TABLE movies(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    directory TEXT,
    release_year INTEGER,
    genre VARCHAR,
    plot TEXT,
    actors TEXT,  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);