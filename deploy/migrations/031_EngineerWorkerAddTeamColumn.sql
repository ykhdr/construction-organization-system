ALTER TABLE engineer_worker
    ADD COLUMN team_id integer NOT NULL REFERENCES engineer_team (id) default 0