SELECT ct.id AS construction_team_id, cp.name AS project
FROM construction_team AS ct
         JOIN work_schedule AS ws ON ct.id = ws.construction_team_id
         JOIN construction_project AS cp ON ct.project_id = cp.id
WHERE ws.fact_start_date > ?
  AND ws.fact_end_date < ?
  AND ws.work_type_id = ?