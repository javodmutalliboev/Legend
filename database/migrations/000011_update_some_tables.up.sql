BEGIN;

ALTER TABLE general_discount RENAME COLUMN title TO title_uz;

ALTER TABLE general_discount 
ADD COLUMN title_ru text NOT NULL UNIQUE DEFAULT 'заголовок',
ADD COLUMN title_en text NOT NULL UNIQUE DEFAULT 'title';

ALTER TABLE general_discount ALTER COLUMN title_uz SET DEFAULT 'sarlavha';

ALTER TABLE goods RENAME COLUMN name TO name_uz;
ALTER TABLE goods ADD COLUMN name_ru text NOT NULL DEFAULT 'имя';
ALTER TABLE goods ADD COLUMN name_en text NOT NULL DEFAULT 'name';

ALTER TABLE goods RENAME COLUMN brand TO brand_uz;
ALTER TABLE goods ADD COLUMN brand_ru text NOT NULL DEFAULT 'бренд';
ALTER TABLE goods ADD COLUMN brand_en text NOT NULL DEFAULT 'brand';

ALTER TABLE goods RENAME COLUMN colors TO colors_uz;
ALTER TABLE goods ADD COLUMN colors_ru text[] NOT NULL DEFAULT '{}';
ALTER TABLE goods ADD COLUMN colors_en text[] NOT NULL DEFAULT '{}';

ALTER TABLE goods RENAME COLUMN description TO description_uz;
ALTER TABLE goods ADD COLUMN description_ru text NOT NULL DEFAULT 'описание';
ALTER TABLE goods ADD COLUMN description_en text NOT NULL DEFAULT 'description';

ALTER TABLE legend_information RENAME COLUMN heading TO heading_uz;
ALTER TABLE legend_information ADD COLUMN heading_ru text NOT NULL DEFAULT 'заголовок';
ALTER TABLE legend_information ADD COLUMN heading_en text NOT NULL DEFAULT 'heading';

ALTER TABLE legend_information RENAME COLUMN description TO description_uz;
ALTER TABLE legend_information ADD COLUMN description_ru text NOT NULL DEFAULT 'описание';
ALTER TABLE legend_information ADD COLUMN description_en text NOT NULL DEFAULT 'description';

ALTER TABLE menu RENAME COLUMN title TO title_uz;
ALTER TABLE menu ADD COLUMN title_ru text NOT NULL DEFAULT 'заголовок';
ALTER TABLE menu ADD COLUMN title_en text NOT NULL DEFAULT 'title';

ALTER TABLE menu_type RENAME COLUMN title TO title_uz;
ALTER TABLE menu_type ADD COLUMN title_ru text UNIQUE;
ALTER TABLE menu_type ADD COLUMN title_en text UNIQUE;

UPDATE menu_type SET title_ru = 'заголовок' || id, title_en = 'title' || id;

ALTER TABLE menu_type ALTER COLUMN title_ru SET NOT NULL;
ALTER TABLE menu_type ALTER COLUMN title_en SET NOT NULL;

COMMIT;