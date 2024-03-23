CREATE TABLE IF NOT EXISTS engineer_worker
(
    skill_level varchar(40) NOT NULL,
    position_id integer     NOT NULL REFERENCES engineer_position (id),
    UNIQUE (id)
) INHERITS (employee)