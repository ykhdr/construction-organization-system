CREATE TABLE IF NOT EXISTS construction_machinery
(
    id         serial PRIMARY KEY,
    name       varchar(150) NOT NULL,
    quantity   integer      NOT NULL CHECK ( quantity > 0 ),
    project_id integer      NOT NULL REFERENCES construction_project (id)
)