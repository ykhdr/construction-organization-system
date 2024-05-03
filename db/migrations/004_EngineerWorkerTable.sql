CREATE TABLE IF NOT EXISTS engineer_worker
(
    position_id integer NOT NULL REFERENCES engineer_position (id),
    team_id     integer NOT NULL REFERENCES engineer_team (id),
    UNIQUE (id)
) INHERITS (employee)