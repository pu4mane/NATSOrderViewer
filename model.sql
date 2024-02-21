CREATE TABLE delivery (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    phone VARCHAR(255),
    zip VARCHAR(255),
    city VARCHAR(255),
    address VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255)
);

CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    transaction VARCHAR(255),
    request_id VARCHAR(255),
    currency VARCHAR(255),
    provider VARCHAR(255),
    amount INTEGER,
    payment_dt INTEGER,
    bank    VARCHAR(255),
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);

CREATE TABLE items(
    id SERIAL PRIMARY KEY,    
    chrt_id INTEGER,
    track_number VARCHAR(255),
    price INTEGER,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INTEGER,
    size VARCHAR(255),
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR(255),
    status INTEGER
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    order_uid varchar(255),
    track_number varchar(255),
    entry varchar(255),
    delivery_id INTEGER REFERENCES delivery(id),
    payment_id INTEGER REFERENCES payment(id),
    locale varchar(255),
    internal_signature varchar(255),
    customer_id varchar(255),
    delivery_service varchar(255),
    shardkey varchar(255),
    sm_id INTEGER,
    date_created varchar(255),
    oof_shard varchar(255)
);