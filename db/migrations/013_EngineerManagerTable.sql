CREATE TABLE IF NOT EXISTS engineer_manager
(
    worker_id integer REFERENCES engineer_worker (id),
    team_id   integer UNIQUE REFERENCES engineer_team (id),
    PRIMARY KEY (worker_id, team_id)
)