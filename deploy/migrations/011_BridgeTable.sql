CREATE TABLE IF NOT EXISTS bridge
(
    span                 integer NOT NULL CHECK ( span > 0 ),
    width                integer NOT NULL CHECK ( width > 0 ),
    traffic_lanes_number integer NOT NULL CHECK ( traffic_lanes_number > 0 )
) INHERITS (construction_project)