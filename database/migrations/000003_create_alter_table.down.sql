BEGIN;

-- Drop menu_type table if exists
DROP TABLE IF EXISTS menu_type;

-- Alter menu table, remove type column and its constraints
ALTER TABLE menu
DROP CONSTRAINT IF EXISTS fk_menu_type,
DROP COLUMN IF EXISTS type;

COMMIT;
