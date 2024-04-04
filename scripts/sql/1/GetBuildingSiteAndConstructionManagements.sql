SELECT cm.name         AS construction_management_name,
       bs.address      AS buidling_site_address,
       cm_manager.name AS management_manager_name,
       bs_manager.name AS building_site_manager_name
FROM building_site AS bs
         JOIN construction_management AS cm ON bs.management_id = cm.id
         LEFT JOIN engineer_worker AS cm_manager ON cm.manager_id = cm_manager.id
         LEFT JOIN engineer_worker AS bs_manager ON bs.manager_id = bs_manager.id
WHERE bs.id != 0
  AND cm.id != 0