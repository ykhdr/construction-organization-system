ALTER TABLE construction_worker
    ADD COLUMN team_id integer NOT NULL REFERENCES construction_team (id) default 0