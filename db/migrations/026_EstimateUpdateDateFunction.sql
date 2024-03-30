CREATE OR REPLACE FUNCTION fn_estimate_update_last_update_date()
    RETURNS trigger AS
$$
BEGIN
    INSERT INTO estimate (id, last_update_date) VALUES (NEW.estimate_id, CURRENT_DATE);
    RETURN NEW;
END
$$ LANGUAGE plpgsql