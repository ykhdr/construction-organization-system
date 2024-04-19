CREATE TRIGGER tr_apartment_house_BI
    BEFORE INSERT ON apartment_house
FOR EACH ROW EXECUTE FUNCTION fn_project_inherits_synchronize_project_id();