CREATE TABLE IF NOT EXISTS construction_worker
(
    is_shift_worker boolean NOT NULL,
    team_id         integer NOT NULL REFERENCES construction_team (id),
    UNIQUE (id)
) INHERITS (employee)