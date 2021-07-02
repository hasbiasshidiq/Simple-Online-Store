
CREATE TABLE public.sellers (
    seller_id character varying(100) NOT NULL PRIMARY KEY,
    seller_name character varying(100) NOT NULL
);

CREATE TABLE public.customers (
    customer_id character varying(100) NOT NULL PRIMARY KEY,
    first_name character varying(100) NOT NULL,
    last_name character varying(100)
);

CREATE TABLE public.products (
    product_id SERIAL PRIMARY KEY,
    product_name character varying(100) NOT NULL,
    category character varying(100) NOT NULL,
    price int NOT NULL
);

CREATE TABLE public.inventory (
    seller_id character varying(100) NOT NULL,
    product_id int NOT NULL,
    quantity int NOT NULL,
    CONSTRAINT unique_inventory UNIQUE (seller_id, product_id),
	CONSTRAINT inventory_seller_id_fk FOREIGN KEY (seller_id) REFERENCES sellers(seller_id),
    CONSTRAINT inventory_product_id_fk FOREIGN KEY (product_id) REFERENCES products(product_id)
);

CREATE TABLE public.orders (
    order_id SERIAL PRIMARY KEY,
    seller_id character varying(100) NOT NULL,
    customer_id character varying(100) NOT NULL,
    order_status character varying(100) NOT NULL,
	order_time timestamptz NOT NULL,
	CONSTRAINT orders_seller_id_fk FOREIGN KEY (seller_id) REFERENCES sellers(seller_id),
    CONSTRAINT orders_customer_id_fk FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

CREATE TABLE public.order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id int NOT NULL,
    product_id int NOT NULL,
    quantity int NOT NULL,
	total_price int NOT NULL,
	CONSTRAINT orders_items_order_id_fk FOREIGN KEY (order_id) REFERENCES orders(order_id),
    CONSTRAINT orders_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products(product_id)
);
