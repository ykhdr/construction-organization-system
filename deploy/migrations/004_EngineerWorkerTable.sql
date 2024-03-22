CREATE TABLE IF NOT EXISTS engineer_worker
(
    skill_level varchar(40) NOT NULL,
    UNIQUE (id)
) INHERITS (employee)