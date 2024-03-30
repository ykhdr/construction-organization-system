CREATE TABLE IF NOT EXISTS engineer_position
(
    id   serial PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE
)