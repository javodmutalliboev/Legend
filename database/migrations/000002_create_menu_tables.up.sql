CREATE TABLE IF NOT EXISTS menu (
    id SERIAL PRIMARY KEY,
    parent_id INTEGER REFERENCES menu(id),
    title TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/*
WITH RECURSIVE menu_hierarchy AS (
    SELECT id, title, parent_id, 1 AS level
    FROM menu
    WHERE parent_id IS NULL

    UNION ALL

    SELECT m.id, m.title, m.parent_id, mh.level + 1
    FROM menu m
    INNER JOIN menu_hierarchy mh ON m.parent_id = mh.id
)
SELECT * FROM menu_hierarchy;
*/