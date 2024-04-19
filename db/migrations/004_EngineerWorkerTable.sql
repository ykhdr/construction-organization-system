CREATE TABLE IF NOT EXISTS engineer_worker
(
    position_id integer     NOT NULL REFERENCES engineer_position (id),
    UNIQUE (id)
) INHERITS (employee)