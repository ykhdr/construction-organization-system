CREATE TABLE IF NOT EXISTS building_site
(
    id            serial PRIMARY KEY,
    address       varchar(150) NOT NULL UNIQUE,
    management_id integer      NOT NULL REFERENCES construction_management (id),
    manager_id    integer      NOT NULL REFERENCES engineer_worker (id)
)