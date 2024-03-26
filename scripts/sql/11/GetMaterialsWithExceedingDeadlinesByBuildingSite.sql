SELECT m.name AS material
FROM building_site AS bs
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN estimate AS e ON cp.id = e.id
         JOIN material_usage AS mu ON e.id = mu.estimate_id
         JOIN material AS m ON mu.material_id = m.id
WHERE bs.id = ?
  AND mu.fact_quantity > mu.plan_quantity