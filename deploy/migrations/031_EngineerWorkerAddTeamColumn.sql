ALTER TABLE engineer_worker
    ADD COLUMN IF NOT EXISTS team_id integer NOT NULL REFERENCES engineer_team (id) default 0