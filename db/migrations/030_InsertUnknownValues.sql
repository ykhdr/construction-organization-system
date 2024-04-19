INSERT INTO engineer_position(id, name)
SELECT 0, 'UNKNOWN'
WHERE NOT EXISTS(SELECT id
                 FROM engineer_position
                 WHERE id = 0);

INSERT INTO building_organization(id, name)
SELECT 0, 'UNKNOWN'
WHERE NOT EXISTS(SELECT id
                 FROM building_organization
                 WHERE id = 0);

INSERT INTO engineer_worker(id, name, surname, patronymic, age, seniority, building_organization_id,
                            position_id)
SELECT 0,
       'UNKNOWN',
       'UNKNOWN',
       '',
       1,
       1,
       0,
       0
WHERE NOT EXISTS(SELECT id
                 FROM engineer_worker
                 WHERE id = 0);

INSERT INTO construction_management(id, name, manager_id)
SELECT 0, 'UNKNOWN', 0
WHERE NOT EXISTS(SELECT id
                 FROM construction_management
                 WHERE id = 0);

INSERT INTO building_site(id, address, management_id, manager_id)
SELECT 0, 'UNKNOWN', 0, 0
WHERE NOT EXISTS(SELECT id
                 FROM building_site
                 WHERE id = 0);

INSERT INTO construction_project(id, name, building_site_id)
SELECT 0, 'UNKNOWN', 0
WHERE NOT EXISTS(SELECT id
                 FROM construction_project
                 WHERE id = 0);

INSERT INTO engineer_team(id, name, project_id)
SELECT 0, 'UNKNOWN', 0
WHERE NOT EXISTS(SELECT id
                 FROM engineer_team
                 WHERE id = 0);

INSERT INTO construction_team(id, name, project_id)
SELECT 0, 'UNKNOWN', 0
WHERE NOT EXISTS(SELECT id
                 FROM construction_team
                 WHERE id = 0);

INSERT INTO construction_worker(id, name, surname, patronymic, age, seniority, building_organization_id,
                                is_shift_worker)
SELECT 0,
       'UNKNOWN',
       'UNKNOWN',
       '',
       1,
       1,
       0,
       false
WHERE NOT EXISTS(SELECT id
                 FROM construction_worker
                 WHERE id = 0);