CREATE TRIGGER tr_bridge_BI
    BEFORE INSERT ON bridge
FOR EACH ROW EXECUTE FUNCTION fn_project_inherits_synchronize_project_id();