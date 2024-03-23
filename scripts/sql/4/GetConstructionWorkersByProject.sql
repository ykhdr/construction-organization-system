SELECT ct.id AS construction_team_id, cw.name AS construction_worker
FROM construction_project AS cp
         JOIN work_schedule AS ws ON cp.id = ws.project_id
         JOIN construction_team AS ct ON ct.id = ws.construction_team_id
         JOIN construction_worker AS cw ON cw.team_id = ct.id
WHERE ws.fact_start_date <= current_date