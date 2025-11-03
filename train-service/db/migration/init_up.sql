CREATE TABLE IF NOT EXISTS trains (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    total_seats INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NOT NULL
);
