CREATE OR REPLACE FUNCTION fn_estimate_update_last_update_date()
    RETURNS trigger AS
$$
BEGIN
    UPDATE estimate SET last_update_date = CURRENT_DATE WHERE id = NEW.estimate_id;
    RETURN NEW;
END
$$ LANGUAGE plpgsql