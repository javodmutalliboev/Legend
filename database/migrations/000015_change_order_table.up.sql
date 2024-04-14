-- Add the total_price column to the order table
ALTER TABLE "order" ADD COLUMN total_price DOUBLE PRECISION NOT NULL DEFAULT 0;

UPDATE "order"
SET total_price = (
    SELECT COALESCE(SUM(order_goods.quantity * (CASE WHEN goods.discount > 0 THEN goods.discount ELSE goods.price END)), 0)
    FROM order_goods
    JOIN goods ON order_goods.goods_id = goods.id
    WHERE order_goods.order_id = "order".id
);