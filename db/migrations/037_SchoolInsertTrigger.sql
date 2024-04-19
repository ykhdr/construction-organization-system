CREATE TRIGGER tr_school_BI
    BEFORE INSERT ON school
FOR EACH ROW EXECUTE FUNCTION fn_project_inherits_synchronize_project_id();