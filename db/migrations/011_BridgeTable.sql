CREATE TABLE IF NOT EXISTS bridge
(
    project_id           integer PRIMARY KEY REFERENCES construction_project (id) ON DELETE CASCADE,
    span                 integer NOT NULL CHECK ( span > 0 ),
    width                integer NOT NULL CHECK ( width > 0 ),
    traffic_lanes_number integer NOT NULL CHECK ( traffic_lanes_number > 0 )
)