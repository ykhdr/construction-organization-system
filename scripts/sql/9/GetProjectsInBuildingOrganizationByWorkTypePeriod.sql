SELECT cp.name AS project
FROM building_organization AS bo
         JOIN construction_contract AS cc ON bo.id = cc.building_organization_id
         JOIN construction_project AS cp ON cc.project_id = cp.id
         JOIN work_schedule AS ws ON cp.id = ws.project_id
WHERE ws.work_type_id = ?
  AND ws.fact_start_date >= ?
  AND ws.fact_end_date <= ?