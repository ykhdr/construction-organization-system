CREATE TABLE IF NOT EXISTS employee
(
    id                       serial PRIMARY KEY,
    name                     varchar(100) NOT NULL,
    surname                  varchar(100) NOT NULL,
    patronymic               varchar(100) NOT NULL,
    age                      integer      NOT NULL CHECK ( age > 0 ),
    seniority                integer      NOT NULL CHECK ( seniority > 0 ),
    building_organization_id integer      NOT NULL REFERENCES building_organization (id)
)