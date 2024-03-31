CREATE TABLE IF NOT EXISTS work_schedule
(
    id                   serial PRIMARY KEY,
    construction_team_id integer NOT NULL REFERENCES construction_team (id),
    work_type_id         integer NOT NULL REFERENCES work_type (id),
    plan_start_date      date    NOT NULL CHECK ( plan_start_date >= to_date('01.01.1900', 'DD-MM-YYYY')),
    plan_end_date        date CHECK ( plan_end_date >= to_date('01.01.1900', 'DD-MM-YYYY') AND
                                      plan_end_date >= plan_start_date),
    fact_start_date      date    NOT NULL CHECK ( fact_start_date >= to_date('01.01.1900', 'DD-MM-YYYY')),
    fact_end_date        date    NOT NULL CHECK ( fact_end_date >= to_date('01.01.1900', 'DD-MM-YYYY') AND
                                                  fact_end_date >= fact_start_date),
    plan_order           integer CHECK ( plan_order > 0 ),
    fact_order           integer NOT NULL CHECK ( fact_order > 0 ),
    project_id           integer REFERENCES construction_project (id),

    UNIQUE (construction_team_id, work_type_id, plan_start_date)
)