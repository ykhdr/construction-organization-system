CREATE TABLE IF NOT EXISTS technical_team
(
    team_id                 integer PRIMARY KEY REFERENCES engineer_team (id) ON DELETE CASCADE,
    is_maintain_machinery boolean NOT NULL
)