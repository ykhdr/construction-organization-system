CREATE TABLE IF NOT EXISTS construction_team
(
    id         serial PRIMARY KEY,
    project_id integer NOT NULL REFERENCES construction_project (id)
)