SELECT cm.name AS machinery_name
FROM construction_project AS cp
         JOIN work_schedule AS ws ON cp.id = ws.project_id
         JOIN construction_machinery_usage AS cmu ON ws.id = cmu.work_schedule_id
         JOIN construction_machinery AS cm ON cp.id = cm.project_id
WHERE ws.fact_start_date >= ? -- start date
  AND ws.fact_end_date <= ?   -- end date