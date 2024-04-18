CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE SCHEMA IF NOT EXISTS productcard;
SET search_path TO productcard;
create table if not exists products (
    id uuid not null default public.uuid_generate_v4(),
    name varchar(255) not null,
    description text,
    price bigint not null check ( price >= 0 )
    );

CREATE SCHEMA IF NOT EXISTS accountservice;
SET search_path TO accountservice;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE user_account (
                              id uuid default public.uuid_generate_v4() PRIMARY KEY,
                              name varchar(100),
                              surname varchar(100),
                              login varchar(100),
                              password varchar(100),
                              FOREIGN KEY (id) REFERENCES user_account (id)
);

CREATE TABLE cart (
                      id uuid default public.uuid_generate_v4() PRIMARY KEY,
                      userID uuid REFERENCES user_account (id)
);

CREATE TABLE product (
                         id uuid PRIMARY KEY,
                         name varchar(100),
                         description varchar(255),
                         price int
);
CREATE TYPE status_enum AS ENUM ('created', 'prepare', 'delivery', 'await', 'received');
CREATE TABLE u_order (
                           id uuid PRIMARY KEY,
                           userID uuid REFERENCES user_account (id),
                           order_status status_enum not null,
                           createdAt timestamp with time zone DEFAULT current_timestamp,
                           updatedAt timestamp with time zone DEFAULT current_timestamp
);


CREATE TABLE product_order (
                               productID uuid REFERENCES product (id),
                               orderID uuid REFERENCES u_order (id),
                               PRIMARY KEY (productID, orderID)
);

CREATE TABLE product_cart (
                             cartID uuid REFERENCES cart (id),
                             productID uuid,
                             PRIMARY KEY (cartID, productID)
);

CREATE SCHEMA IF NOT EXISTS order_service;
SET search_path TO order_service;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE status_enum AS ENUM ('created', 'prepare', 'delivery', 'await', 'received');

create table if not exists orders (
    id uuid default public.uuid_generate_v4() PRIMARY KEY,
    order_status status_enum not null default 'created',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

Insert into accountservice.product values ('21167535-5b55-4cf7-a92d-4a24317932ab', 'myName', 'myDiscr', 20) on conflict (id) do update set name = EXCLUDED.name, description = EXCLUDED.description, price = EXCLUDED.price;
Insert into accountservice.product values ('21167535-5b55-4cf7-a92d-4a24317932ac', 'myName', 'myDiscr', 31) on conflict (id) do update set name = EXCLUDED.name, description = EXCLUDED.description, price = EXCLUDED.price;
INSERT INTO accountservice.user_account (id, Name, Surname, login, password) VALUES ('31167546-5b55-4cf7-a92d-4a24317934bb', 'John', 'Doe', 'johndoe', 'password123');
INSERT INTO accountservice.cart values ('ead7d5ca-f982-43d6-a8f0-2534f1df07d6', '31167546-5b55-4cf7-a92d-4a24317934bb');
SELECT * from accountservice.cart;
INSERT INTO accountservice.product_cart values ('ead7d5ca-f982-43d6-a8f0-2534f1df07d6', '21167535-5b55-4cf7-a92d-4a24317932ab');
INSERT INTO accountservice.product_cart values ('ead7d5ca-f982-43d6-a8f0-2534f1df07d6', '21167535-5b55-4cf7-a92d-4a24317932ac');
SELECT * FROM accountservice.product_cart;
INSERT INTO productcard.products (name, description, price) values ('superProduct1', 'Крутой продукт', 12);
INSERT INTO productcard.products (name, description, price) values ('superProduct2', 'Еще один крутой продукт', 15);
Select * from productcard.products;
Select * from order_service.orders;
SELECT p.id, p.name, p.description, p.price FROM accountservice.product p JOIN accountservice.product_cart pc ON p.id = pc.productID JOIN accountservice.cart c ON pc.cartID = c.id JOIN accountservice.user_account ua ON c.userID = ua.id WHERE ua.id = '31167546-5b55-4cf7-a92d-4a24317934bb';

