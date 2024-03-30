CREATE OR REPLACE TRIGGER tr_estimate_BU_BI
    BEFORE INSERT OR UPDATE OF fact_quantity
    ON material_usage
    FOR EACH ROW
EXECUTE PROCEDURE fn_estimate_update_last_update_date()