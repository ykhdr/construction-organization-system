SELECT cp.name AS project
FROM construction_management AS cm
         JOIN building_site AS bs ON cm.id = bs.management_id
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN work_schedule AS ws ON cp.id = ws.project_id
WHERE ws.work_type_id = ?
AND ws.fact_start_date >= ?
AND ws.fact_end_date <= ?