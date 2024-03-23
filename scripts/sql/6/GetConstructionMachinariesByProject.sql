SELECT m.name
FROM construction_project AS cp
         JOIN construction_machinery AS m ON cp.id = m.project_id
WHERE cp.name = ?
  AND cp.id != 0