CREATE TABLE IF NOT EXISTS bookings (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    schedule_id BIGINT NOT NULL,
    booking_code VARCHAR(20) UNIQUE NOT NULL,
    seat_count INTEGER NOT NULL,
    total_price DECIMAL(10,2) DEFAULT 0,
    status INTEGER DEFAULT 1, -- 1=pending, 2=success, 3=failed, 4=expired
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_schedule_id ON bookings(schedule_id);
CREATE INDEX idx_bookings_booking_code ON bookings(booking_code);
CREATE INDEX idx_bookings_status ON bookings(status);
CREATE INDEX idx_bookings_expires_at ON bookings(expires_at);
CREATE INDEX idx_bookings_deleted_at ON bookings(deleted_at);