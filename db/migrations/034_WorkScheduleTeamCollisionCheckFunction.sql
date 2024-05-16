CREATE OR REPLACE FUNCTION fn_work_schedule_team_collision_check()
    RETURNS trigger AS
$$
BEGIN
    IF EXISTS(SELECT 1
              FROM work_schedule
              WHERE construction_team_id = NEW.construction_team_id
                AND (
                  (NEW.plan_start_date BETWEEN plan_start_date AND plan_end_date) OR
                  (NEW.plan_end_date BETWEEN plan_start_date AND plan_end_date) OR
                  (plan_start_date BETWEEN NEW.plan_start_date AND NEW.plan_end_date) OR
                  (plan_end_date BETWEEN NEW.plan_start_date AND NEW.plan_end_date) OR
                  (NEW.fact_start_date IS NOT NULL AND
                   NEW.fact_start_date BETWEEN fact_start_date AND fact_end_date) OR
                  (NEW.fact_end_date IS NOT NULL AND
                   NEW.fact_end_date BETWEEN fact_start_date AND fact_end_date) OR
                  (fact_start_date IS NOT NULL AND
                   fact_start_date BETWEEN NEW.fact_start_date AND NEW.fact_end_date) OR
                  (fact_end_date IS NOT NULL AND
                   fact_end_date BETWEEN NEW.fact_start_date AND NEW.fact_end_date)
                  ))
    THEN
        RAISE EXCEPTION 'Team has a work in schedule at same time';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;