DROP TABLE IF EXISTS coupon;

CREATE TABLE coupon
(
    id SERIAL PRIMARY KEY,
    token VARCHAR(10) UNIQUE,
    email VARCHAR(120),
    address VARCHAR(120),
    card_id INT NOT NULL
);