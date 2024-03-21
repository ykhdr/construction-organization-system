CREATE TABLE IF NOT EXISTS master_team
(
    design_experience_years integer NOT NULL CHECK ( design_experience_years > 0 )
) INHERITS (engineer_team)