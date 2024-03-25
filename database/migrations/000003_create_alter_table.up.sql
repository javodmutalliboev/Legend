BEGIN;

-- Create menu_type table
CREATE TABLE IF NOT EXISTS menu_type (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Insert two new menu_types
INSERT INTO menu_type (title) VALUES ('Jamoaviy kiyimlar'), ('Korporativ jamoa kiyimlari');

-- Alter menu table
ALTER TABLE menu ADD COLUMN type INTEGER;

-- Set default value for 'type' to the id of the menu_type with the smallest id
UPDATE menu SET type = (SELECT id FROM menu_type ORDER BY id ASC LIMIT 1);

-- Now that all rows have a 'type' value, add the NOT NULL constraint
ALTER TABLE menu ALTER COLUMN type SET NOT NULL;

-- Add foreign key constraint to 'type' referencing menu_type.id
ALTER TABLE menu ADD CONSTRAINT fk_menu_type FOREIGN KEY (type) REFERENCES menu_type(id);

COMMIT;