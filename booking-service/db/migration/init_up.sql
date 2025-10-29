CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    schedule_id INT NOT NULL,
    quantity INT NOT NULL,
    set_seats TEXT,
    status INT NOT NULL DEFAULT 1,
    expired_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);
