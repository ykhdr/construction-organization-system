CREATE TABLE IF NOT EXISTS school
(
    classroom_count integer NOT NULL,
    floors          integer NOT NULL CHECK ( floors > 0 )
) INHERITS (construction_project)