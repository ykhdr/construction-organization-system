SELECT cm.name
FROM construction_project AS cp
         JOIN work_schedule AS ws ON cp.id = ws.project_id
         JOIN work_schedule_construction_machinery AS wscm ON ws.id = wscm.work_schedule_id
         JOIN construction_machinery AS cm ON cp.id = cm.project_id
WHERE ws.fact_start_date >= ? -- start date
  AND ws.fact_end_date <= ?   -- end date