CREATE TABLE IF NOT EXISTS school
(
    project_id      integer PRIMARY KEY REFERENCES construction_project (id) ON DELETE CASCADE,
    classroom_count integer NOT NULL,
    floors          integer NOT NULL CHECK ( floors > 0 )
)