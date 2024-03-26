SELECT wt.name AS work_type, cp.name AS project
FROM construction_team AS ct
         JOIN work_schedule AS ws ON ct.id = ws.construction_team_id
         JOIN construction_project AS cp ON ws.project_id = cp.id
         JOIN work_type AS wt ON ws.work_type_id = wt.id
WHERE ct.id = ?
  AND ws.fact_start_date > ?
  AND ws.fact_end_date < ?