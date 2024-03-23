CREATE TABLE IF NOT EXISTS work_schedule_construction_machinery
(
    work_schedule_id integer REFERENCES work_schedule (id),
    machinery_id     integer REFERENCES construction_machinery (id),
    quantity         integer NOT NULL
)