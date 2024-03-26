SELECT m.name AS material
FROM construction_management AS cm
         JOIN building_site AS bs ON cm.id = bs.management_id
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN estimate AS e ON cp.id = e.id
         JOIN material_usage AS mu ON e.id = mu.estimate_id
         JOIN material AS m ON mu.material_id = m.id
WHERE cm.id = ?
  AND mu.fact_quantity > mu.plan_quantity