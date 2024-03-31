CREATE TABLE IF NOT EXISTS construction_team
(
    id         serial PRIMARY KEY,
    name       varchar(100) NOT NULL,
    project_id integer      NOT NULL REFERENCES construction_project (id),
    UNIQUE (name, project_id)
)