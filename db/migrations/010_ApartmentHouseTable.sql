CREATE TABLE IF NOT EXISTS apartment_house
(
    floors integer NOT NULL CHECK ( floors > 0 ),
    UNIQUE (id)
) INHERITS (construction_project)