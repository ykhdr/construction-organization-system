SELECT cp.name            AS project_name,
       m.name             AS material_name,
       m.cost             AS material_cost,
       mu.plan_quantity   AS plan_quantity,
       mu.fact_quantity   AS fact_quantity,
       e.last_update_date AS estimate_last_update_date
FROM construction_project AS cp
         JOIN estimate AS e ON cp.id = e.id
         JOIN material_usage AS mu ON e.id = mu.estimate_id
         JOIN material AS m ON mu.material_id = m.id
WHERE cp.id = ?
  AND cp.id != 0