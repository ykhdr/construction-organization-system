CREATE TABLE IF NOT EXISTS report
(
    id                   serial PRIMARY KEY,
    project_id           integer NOT NULL REFERENCES construction_project (id),
    report_creation_date date    NOT NULL CHECK ( report_creation_date >= to_date('01.01.1900', 'DD-MM-YYYY')),
    report_file          json    NOT NULL
)