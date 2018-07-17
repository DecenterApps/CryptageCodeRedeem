DROP TABLE IF EXISTS coupon;
DROP TABLE IF EXISTS card;
DROP TABLE IF EXISTS "user";

CREATE TABLE "user"
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120),
    address VARCHAR(120)
);

CREATE TABLE card
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120)
);

CREATE TABLE coupon
(
    id SERIAL PRIMARY KEY,
    token VARCHAR(10) UNIQUE,
    user_id INT REFERENCES "user" (id),
    card_id INT REFERENCES card (id)
);