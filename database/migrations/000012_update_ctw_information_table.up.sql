ALTER TABLE ctw_information RENAME COLUMN heading TO heading_uz;
ALTER TABLE ctw_information ADD COLUMN heading_ru text NOT NULL DEFAULT 'заголовок';
ALTER TABLE ctw_information ADD COLUMN heading_en text NOT NULL DEFAULT 'heading';

ALTER TABLE ctw_information RENAME COLUMN description TO description_uz;
ALTER TABLE ctw_information ADD COLUMN description_ru text NOT NULL DEFAULT 'описание';
ALTER TABLE ctw_information ADD COLUMN description_en text NOT NULL DEFAULT 'description';