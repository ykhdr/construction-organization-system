SELECT cm.name AS construction_management_name, bs.address AS buidling_site_address, ew.name AS manager_name
FROM building_site AS bs
         JOIN construction_management AS cm ON bs.management_id = cm.id
         JOIN engineer_worker AS ew ON ew.id = bs.id
WHERE bs.id != 0
  AND cm.id != 0