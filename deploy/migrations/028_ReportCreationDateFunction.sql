CREATE OR REPLACE FUNCTION fn_report_set_creation_date()
    RETURNS trigger AS
$$
BEGIN
    NEW.report_creation_date := current_date;
    RETURN NEW;
END
$$ LANGUAGE plpgsql