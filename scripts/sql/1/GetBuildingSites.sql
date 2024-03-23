SELECT bs.address AS buidling_site_address, ew.name AS manager_name
FROM building_site AS bs
         JOIN engineer_worker AS ew ON ew.id = bs.id
WHERE bs.id != 0