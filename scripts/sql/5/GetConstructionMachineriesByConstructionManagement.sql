SELECT m.name
FROM construction_management AS cm
         JOIN building_site AS bs ON cm.id = bs.management_id
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN construction_machinery AS m ON cp.id = m.project_id
WHERE cm.name = ?
  AND cm.id != 0