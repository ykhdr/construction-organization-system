SELECT wt.name AS work_type
FROM building_site AS bs
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN work_schedule AS ws ON cp.id = ws.project_id
         JOIN work_type AS wt ON ws.work_type_id = wt.id
WHERE bs.id = ?
  AND ws.fact_end_date > ws.plan_end_date