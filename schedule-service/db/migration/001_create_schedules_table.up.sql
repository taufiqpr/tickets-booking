CREATE TABLE IF NOT EXISTS schedules (
    id BIGSERIAL PRIMARY KEY,
    train_id BIGINT NOT NULL,
    origin VARCHAR(100) NOT NULL,
    destination VARCHAR(100) NOT NULL,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    available_seats INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_schedules_train_id ON schedules(train_id);
CREATE INDEX idx_schedules_origin_destination ON schedules(origin, destination);
CREATE INDEX idx_schedules_departure_time ON schedules(departure_time);
CREATE INDEX idx_schedules_deleted_at ON schedules(deleted_at);