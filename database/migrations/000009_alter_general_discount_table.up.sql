ALTER TABLE general_discount
ADD COLUMN title TEXT NOT NULL UNIQUE DEFAULT 'title';