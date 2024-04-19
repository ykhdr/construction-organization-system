CREATE OR REPLACE FUNCTION fn_project_inherits_synchronize_project_id()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO construction_project (id, name, building_site_id)
    VALUES (NEW.id, NEW.name, NEW.building_site_id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;