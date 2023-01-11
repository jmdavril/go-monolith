CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE product(
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    sku VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    price NUMERIC(6,2) NOT NULL
);
CREATE INDEX product_sku_index ON product (sku);

CREATE TABLE customer(
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE
);

CREATE TABLE shop_order(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    total_spent NUMERIC(6,2) NOT NULL,
    customer_id uuid NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);
CREATE INDEX shop_order_customer_index ON shop_order (customer_id);

CREATE TABLE line_item(
    order_id uuid NOT NULL,
    line_index INT NOT NULL,
    sku VARCHAR NOT NULL,
    quantity INT NOT NULL,
    unit_price NUMERIC(6,2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES shop_order(id),
    PRIMARY KEY(order_id, line_index)
);

CREATE TABLE product_sales(
    sku VARCHAR PRIMARY KEY,
    quantity INT NOT NULL,
    total_sales NUMERIC(6,2) NOT NULL
);
CREATE INDEX product_sales_quantity_index ON product_sales (quantity DESC);
CREATE INDEX product_sales_sales_index ON product_sales (total_sales DESC);
