Insert into accountservice.product values ('21167535-5b55-4cf7-a92d-4a24317932ab', 'myName', 'myDiscr', 20) on conflict (id) do update set name = EXCLUDED.name, description = EXCLUDED.description, price = EXCLUDED.price;
Insert into accountservice.product values ('21167535-5b55-4cf7-a92d-4a24317932ac', 'myName', 'myDiscr', 31) on conflict (id) do update set name = EXCLUDED.name, description = EXCLUDED.description, price = EXCLUDED.price;
INSERT INTO accountservice.user_account (UserID, Name, Surname, login, password) VALUES (1, 'John', 'Doe', 'johndoe', 'password123');
INSERT INTO accountservice.cart values (1, 1);
SELECT * from accountservice.cart;
INSERT INTO accountservice.product_cart values (1, '21167535-5b55-4cf7-a92d-4a24317932ab');
INSERT INTO accountservice.product_cart values (1, '21167535-5b55-4cf7-a92d-4a24317932ac');
SELECT * FROM accountservice.product_cart;
INSERT INTO productcard.products (name, description, price) values ('superProduct2', 'Еще один крутой продукт', 15);
Select * from productcard.products;
Select * from order_service.orders;
SELECT p.id, p.name, p.description, p.price FROM accountservice.product p JOIN accountservice.product_cart pc ON p.id = pc.productID JOIN accountservice.cart c ON pc.cartID = c.cartID JOIN accountservice.user_account ua ON c.userID = ua.userID WHERE ua.userID = 1;