CREATE TABLE IF NOT EXISTS construction_worker
(
    is_shift_worker boolean NOT NULL,
    UNIQUE (id)
) INHERITS (employee)