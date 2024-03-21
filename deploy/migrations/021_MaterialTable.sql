CREATE TABLE IF NOT EXISTS material
(
    id   serial PRIMARY KEY,
    name varchar(150) NOT NULL,
    cost integer      NOT NULL CHECK ( cost > 0 )
)