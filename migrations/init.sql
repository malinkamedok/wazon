CREATE SCHEMA IF NOT EXISTS productcard;
SET search_path TO productcard;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists products (
    id uuid not null default uuid_generate_v4(),
    name varchar(255) not null,
    description text,
    price bigint not null check ( price >= 0 )
    );

CREATE SCHEMA IF NOT EXISTS accountservice;
SET search_path TO accountservice;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE user_account (
                              userID SERIAL PRIMARY KEY,
                              name varchar(100),
                              surname varchar(100),
                              login varchar(100),
                              password varchar(100),
                              FOREIGN KEY (userID) REFERENCES user_account (userID)
);

CREATE TABLE cart (
                      cartID SERIAL PRIMARY KEY,
                      userID int REFERENCES user_account (userID)
);

CREATE TABLE product (
                         id uuid PRIMARY KEY,
                         name varchar(100),
                         description varchar(255),
                         price int
);

CREATE TABLE u_order (
                           orderID SERIAL PRIMARY KEY,
                           userID int REFERENCES user_account (userID),
                           orderState varchar(100),
                           createdAt timestamp with time zone DEFAULT current_timestamp,
                           updatedAt timestamp with time zone DEFAULT current_timestamp
);

CREATE TABLE product_order (
                              productOrderID SERIAL PRIMARY KEY,
                              productID uuid REFERENCES product (id),
                              orderID int REFERENCES u_order (orderID)
);

CREATE TABLE product_cart (
                             cartID SERIAL REFERENCES cart (cartID),
                             productID uuid REFERENCES product (id),
                             PRIMARY KEY (cartID, productID)
);


CREATE TYPE status_enum AS ENUM ('created', 'prepare', 'delivery', 'await', 'received');

create table if not exists orders (
    id uuid not null default uuid_generate_v4() unique,
    order_status status_enum not null default 'created',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);