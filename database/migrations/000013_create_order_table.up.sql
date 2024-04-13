CREATE TABLE IF NOT EXISTS "order" (
    id BIGSERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    customer_surname VARCHAR(255) NOT NULL,
    customer_region VARCHAR(255) NOT NULL,
    customer_district VARCHAR(255) NOT NULL,
    customer_address VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(255) NOT NULL,
    customer_phone2 VARCHAR(255) NOT NULL,
    canceled BOOLEAN NOT NULL DEFAULT FALSE,
    delivered BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP  
);

CREATE TABLE IF NOT EXISTS order_goods (
    goods_id BIGINT NOT NULL REFERENCES "goods" (id),
    order_id BIGINT NOT NULL REFERENCES "order" (id),
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);