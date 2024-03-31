CREATE TABLE IF NOT EXISTS construction_machinery_usage
(
    work_schedule_id integer REFERENCES work_schedule (id),
    machinery_id     integer REFERENCES construction_machinery (id),
    quantity         integer NOT NULL,
    PRIMARY KEY (work_schedule_id, machinery_id)
)