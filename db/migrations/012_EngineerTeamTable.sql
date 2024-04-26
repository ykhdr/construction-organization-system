CREATE TABLE IF NOT EXISTS engineer_team
(
    id         serial PRIMARY KEY,
    name       varchar(100) NOT NULL,
    project_id integer      NOT NULL REFERENCES construction_project (id),
    type       varchar(30) NOT NULL,
    UNIQUE (name, project_id)
)