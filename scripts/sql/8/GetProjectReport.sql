SELECT report.report_file AS report
FROM report
WHERE project_id = ?
  AND project_id != 0