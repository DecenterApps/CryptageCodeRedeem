DROP TABLE IF EXISTS coupon;
DROP TABLE IF EXISTS card;
DROP TABLE IF EXISTS "user";

CREATE TABLE "user"
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(120),
    address VARCHAR(120)
);

CREATE TABLE coupon
(
    id SERIAL PRIMARY KEY,
    token VARCHAR(10) UNIQUE,
    card_id INT NOT NULL,
    user_id INT REFERENCES "user" (id)
);