SELECT cs.name AS buidling_site_name, ew.name AS manager_name
FROM building_site AS cs
         JOIN engineer_worker AS ew ON ew.id = cs.id
