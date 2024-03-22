CREATE TABLE IF NOT EXISTS engineer_worker
(
    is_shift_worker boolean NOT NULL,
    UNIQUE (id)
) INHERITS (employee)