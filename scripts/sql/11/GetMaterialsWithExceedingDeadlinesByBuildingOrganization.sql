SELECT m.name AS material
FROM building_organization AS bo
         JOIN construction_contract AS cc ON bo.id = cc.building_organization_id
         JOIN construction_project AS cp ON cc.project_id = cp.id
         JOIN estimate AS e ON cp.id = e.id
         JOIN material_usage AS mu ON e.id = mu.estimate_id
         JOIN material AS m ON mu.material_id = m.id
WHERE bo.id = ?
  AND mu.fact_quantity > mu.plan_quantity