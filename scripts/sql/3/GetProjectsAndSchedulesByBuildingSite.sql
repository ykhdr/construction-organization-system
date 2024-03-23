SELECT cp.name            AS project_name,
       wt.name            AS work_type,
       ws.plan_start_date AS plan_start_date,
       ws.plan_end_date   AS plan_end_date,
       ws.fact_start_date AS fact_start_date,
       ws.fact_end_date   AS fact_end_date,
       ws.plan_order      AS plan_order,
       ws.fact_order      AS fact_order
FROM building_site AS bs
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN work_schedule AS ws ON cp.id = ws.project_id
         JOIN work_type AS wt ON ws.work_type_id = wt.id
WHERE bs.id = ?
  AND bs.id != 0