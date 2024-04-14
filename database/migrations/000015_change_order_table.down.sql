ALTER TABLE "order" DROP COLUMN total_price;

DROP TRIGGER IF EXISTS update_total_price ON "order";

DROP FUNCTION IF EXISTS calculate_total_price();