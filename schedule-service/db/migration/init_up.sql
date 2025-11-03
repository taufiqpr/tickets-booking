CREATE TABLE IF NOT EXISTS schedules (
    id SERIAL PRIMARY KEY,
    train_id INT NOT NULL,
    origin VARCHAR(50) NOT NULL,
    destination VARCHAR(50) NOT NULL,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    available_seats INT NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NOT NULL
);
