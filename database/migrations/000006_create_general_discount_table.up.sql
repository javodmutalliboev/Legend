CREATE TABLE IF NOT EXISTS general_discount (
    id serial PRIMARY KEY,
    menu_type integer NOT NULL REFERENCES menu_type(id),
    value double precision NOT NULL DEFAULT 0,
    unit text NOT NULL DEFAULT '%',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);