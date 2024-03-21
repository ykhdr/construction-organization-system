CREATE TABLE IF NOT EXISTS construction_manager
(
    worker_id integer REFERENCES construction_worker (id),
    team_id   integer UNIQUE REFERENCES construction_team (id),
    PRIMARY KEY (worker_id, team_id)
)