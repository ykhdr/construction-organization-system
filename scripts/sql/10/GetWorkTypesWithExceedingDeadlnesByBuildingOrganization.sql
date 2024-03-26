SELECT wt.name AS work_type
FROM building_organization AS bo
         JOIN construction_contract AS cc ON bo.id = cc.building_organization_id
         JOIN construction_project AS cp ON cc.project_id = cp.id
         JOIN work_schedule AS ws ON cp.id = ws.project_id
         JOIN work_type AS wt ON ws.work_type_id = wt.id
WHERE bo.id = ?
  AND ws.fact_end_date > ws.plan_end_date