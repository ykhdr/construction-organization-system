CREATE TABLE IF NOT EXISTS construction_machinery
(
    id         serial PRIMARY KEY,
    name       varchar(150) NOT NULL,
    project_id integer      NOT NULL REFERENCES construction_project (id)
)