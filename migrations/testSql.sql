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

