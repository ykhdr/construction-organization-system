CREATE TABLE IF NOT EXISTS master_team
(
    team_id                 integer PRIMARY KEY REFERENCES engineer_team (id) ON DELETE CASCADE,
    design_experience_years integer NOT NULL CHECK ( design_experience_years > 0 )
)