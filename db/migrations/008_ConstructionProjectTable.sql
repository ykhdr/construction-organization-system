CREATE TABLE IF NOT EXISTS construction_project
(
    id               serial PRIMARY KEY,
    name             varchar(150) NOT NULL,
    building_site_id integer      NOT NULL REFERENCES building_site (id),
    type             varchar(30)  NOT NULL
)