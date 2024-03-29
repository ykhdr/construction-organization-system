CREATE TABLE IF NOT EXISTS construction_contract
(
    id                       serial PRIMARY KEY,
    building_organization_id integer      NOT NULL REFERENCES building_organization (id),
    customer_organization_id integer      NOT NULL REFERENCES customer_organization (id),
    name                     varchar(150) NOT NULL,
    project_id               integer      NOT NULL UNIQUE REFERENCES construction_project (id),
    signing_date             date         NOT NULL CHECK ( signing_date >= to_date('01.01.1900', 'DD-MM-YYYY'))
)