DROP VIEW all_operations

CREATE VIEW all_operations AS SELECT 
    c.id
    , unit_name
    , type_name
    , plant_name
    , description
    , author
    , plant_id
FROM units AS c
INNER JOIN types AS a
    ON type_id = a.id
INNER JOIN plants AS b
    ON plant_id = b.id
WHERE type_id = 1

SELECT * FROM all_operations