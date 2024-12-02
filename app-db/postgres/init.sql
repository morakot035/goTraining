CREATE SCHEMA app_data;

CREATE TABLE IF NOT EXISTS app_data.restaurants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    menu JSONB
);

CREATE TABLE IF NOT EXISTS app_data.riders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS app_data.orders (
    id SERIAL PRIMARY KEY,
    restaurantId VARCHAR(10)
    riderId VARCHAR(10)
    items JSONB
    status VARCHAR(20)
);