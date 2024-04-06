SELECT ew.name AS engineer_name, ep.name AS engineer_position
FROM building_site AS bs
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN engineer_team AS et ON cp.id = et.project_id
         JOIN engineer_worker AS ew ON et.id = ew.team_id
         JOIN engineer_position AS ep ON ew.position_id = ep.id
WHERE bs.id = ?
  AND bs.id != 0