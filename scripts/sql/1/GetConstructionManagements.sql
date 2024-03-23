SELECT cm.name AS construction_management_name, ew.name AS manager_name
FROM construction_management AS cm
         JOIN engineer_worker AS ew ON ew.id = cm.id
WHERE cm.id != 0
