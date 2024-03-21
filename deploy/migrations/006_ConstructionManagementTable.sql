CREATE TABLE IF NOT EXISTS construction_management
(
    id         serial PRIMARY KEY,
    name       varchar(100) NOT NULL,
    manager_id integer      NOT NULL REFERENCES engineer_worker (id)
)