CREATE TABLE IF NOT EXISTS trainers (
    trainer_id SERIAL PRIMARY KEY,
    gym_id INT,
    trainer_name VARCHAR,
    speciality VARCHAR
);

CREATE TABLE IF NOT EXISTS current_bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id BIGINT,
    trainer_id INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    activity VARCHAR,
    FOREIGN KEY (trainer_id) REFERENCES trainers(trainer_id)
);

CREATE TABLE IF NOT EXISTS available_bookings (
    booking_id SERIAL PRIMARY KEY,
    trainer_id INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    activity VARCHAR,
    FOREIGN KEY (trainer_id) REFERENCES trainers(trainer_id)
);
