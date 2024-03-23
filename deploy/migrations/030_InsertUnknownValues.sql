INSERT INTO engineer_position(id, name)
VALUES (0, 'Unknown');

INSERT INTO building_organization(id, name)
VALUES (0, 'Unknown');

INSERT INTO engineer_worker(id, name, surname, patronymic, age, seniority, building_organization_id, skill_level,
                            position_id)
VALUES (0, 'Unknown', 'Unknown', null, 1, 1, 0, 'Unknown', 0);

INSERT INTO construction_management(id, name, manager_id)
VALUES (0, 'Unknown', 0);

INSERT INTO building_site(id, address, management_id, manager_id)
VALUES (0, 'Unknown', 0, 0);

INSERT INTO construction_project(id, name, building_site_id)
VALUES (0, 'Unknown', 0);

INSERT INTO engineer_team(id, project_id)
VALUES (0, 0);

INSERT INTO construction_worker(id, name, surname, patronymic, age, seniority, building_organization_id,
                                is_shift_worker)
VALUES (0, 'Unknown', 'Unknown', null, 1, 1, 0, false);