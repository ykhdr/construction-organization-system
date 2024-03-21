CREATE TABLE IF NOT EXISTS work_type
(
    id   integer PRIMARY KEY,
    name varchar(100) UNIQUE NOT NULL
)