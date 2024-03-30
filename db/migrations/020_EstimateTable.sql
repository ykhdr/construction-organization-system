CREATE TABLE IF NOT EXISTS estimate
(
    id               integer PRIMARY KEY REFERENCES construction_project (id),
    last_update_date date NOT NULL CHECK ( last_update_date >= to_date('01.01.1900', 'DD-MM-YYYY') )
)