SELECT ew.name AS engineer_name, ep.name AS engineer_position
FROM construction_management AS cm
         JOIN building_site AS bs ON cm.id = bs.management_id
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN engineer_team AS et ON cp.id = et.project_id
         JOIN engineer_worker AS ew ON et.id = ew.team_id
         JOIN engineer_position AS ep ON ew.position_id = ep.id
WHERE building_site_id = ?
  AND building_site_id != 0