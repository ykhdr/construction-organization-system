CREATE OR REPLACE FUNCTION fn_work_schedule_team_collision_check()
    RETURNS trigger AS
$$
BEGIN
    IF EXISTS(SELECT *
              FROM work_schedule
              WHERE construction_team_id = NEW.construction_team_id
                AND (NEW.plan_start_date BETWEEN plan_start_date AND plan_end_date OR
                     NEW.fact_start_date BETWEEN fact_start_date AND fact_end_date))
    THEN
        RAISE WARNING 'Team has a work in schedule at same time';
    END IF;
END;
$$ language plpgsql