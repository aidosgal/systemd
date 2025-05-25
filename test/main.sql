CREATE TABLE IF NOT EXISTS users (
    id PRIMARY KEY SERIAL,
    name VARCHAR(255)
);

CREATE TYPE ride_status AS ENUM('requested','in_progress','completed', 'cancelled');

CREATE TABLE IF NOT EXISTS rides (
    id PRIMARY KEY SERIAL,
    passenger_id INTEGER NOT NULL REFERENCES users(id),
    driver_id INTEGER NOT NULL REFERENCES users(id),
    from_city VARCHAR(255) NOT NULL,
    to_city VARCHAR(255) NOT NULL,
    status ride_status NOT NULL,
    started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ended_at TIMESTAMP NULL
);

CREATE TYPE payment_status AS ENUM('pending','paid','failed');

CREATE TABLE IF NOT EXISTS payments (
    id PRIMARY KEY SERIAL,
    ride_id INTEGER NOT NULL REFERENCES rides(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    amount NUMERIC(24, 8) NOT NULL,
    status payment_status NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(ride_id, id)
);

SELECT * FROM rides
WHERE
    started_at BETWEEN NOW() - INTERVAL 7 days AND NOW()
    AND (passenger_id = 42 OR driver_id = 42);

SELECT SUM(amount)
FROM payments p
LEFT JOIN rides r ON r.id = p.ride_id
WHERE p.to_city = "Almaty" AND p.status = "paid";

SELECT *
FROM users u
LEFT JOIN rides r ON r.driver_id = u.id
WHERE r.status = "completed"
ORDER BY COUNT(u.id) DESC
LIMIT 10;

