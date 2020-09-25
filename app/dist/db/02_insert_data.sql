--
-- Script para insertar datos de configuracion en el sistema
--

--
-- USERS (usuarios)
--                                          
insert into users (username, name, surname, email, password) values ('cristian', 'cristian', 'rojas', 'cristian@cr.net', 'd8c15d56990027dfe431d9b8430c2797ea7aaa7f61b586f09a755c3de57f32871682db151b359c42847739b919b558bfbe365ee6d1c0b30fa0db0da395b162a7');
insert into users (username, name, surname, email, password) values ('doris', 'doris', 'sandoval', 'doris@cr.net', 'd8c15d56990027dfe431d9b8430c2797ea7aaa7f61b586f09a755c3de57f32871682db151b359c42847739b919b558bfbe365ee6d1c0b30fa0db0da395b162a7');

--
--CATEGORIES (categorias) Y CATEGORIES_RELATIONS (relaciones entre categorias)
--
--------------------------------------------------------------------------------------------------------------------------------
--Notebooks y PC (principal)
insert into categories (id, main, name, icon_name) values (1,true,'Notebooks y PC','desktop_mac');

--Notebooks y PC -- Notebooks
insert into categories (id, main, name, icon_name) values (2,false,'Notebooks','');

--Notebooks y PC -- Impresoras y Accesorios
insert into categories (id, main, name, icon_name) values (3,false,'Impresoras y Accesorios','');

--Notebooks y PC -- PC de Escritorio
insert into categories (id, main, name, icon_name) values (4,false,'PC de Escritorio','');


--Notebooks y PC -- Notebooks
insert into categories_relations (id_main, id_child) values (1,2);
--Celulares y Smartwatch -- Impresoras y Accesorios
insert into categories_relations (id_main, id_child) values (1,3);
--Celulares y Smartwatch -- PC de Escritorio
insert into categories_relations (id_main, id_child) values (1,4);

--------------------------------------------------------------------------------------------------------------------------------
-- Celulares y Smartwatch
insert into categories (id, main, name, icon_name) values (5,true,'Celulares y Smartwatch','devices_other');
-- Celulares y Smartwatch -- Smartwatch
insert into categories (id, main, name, icon_name) values (6,false,'Smartwatch','');
-- Celulares y Smartwatch -- Celulares
insert into categories (id, main, name, icon_name) values (7,false,'Celulares','');

--Celulares y Smartwatch -- Smartwatch
insert into categories_relations (id_main, id_child) values (5,6);
--Celulares y Smartwatch -- Celulares
insert into categories_relations (id_main, id_child) values (5,7);

--------------------------------------------------------------------------------------------------------------------------------

--Hogar y Electrodomesticos
insert into categories (id, main, name, icon_name) values (8,true,'Hogar y Electrodomesticos','electrical_services');
--Hogar y Electrodomesticos -- Electrodomesticos y aires
insert into categories (id, main, name, icon_name) values (9,false,'Electrodomesticos y aires','');

--Hogar y Electrodomesticos -- Electrodomesticos y aires
insert into categories_relations (id_main, id_child) values (8,9);

--------------------------------------------------------------------------------------------------------------------------------
--Accesorios
insert into categories (id, main, name, icon_name) values (10,true,'Accesorios','headset');

--Accesorios -- Audio
insert into categories (id, main, name, icon_name) values (11,false,'Audio','');

--Accesorios -- Auriculares
insert into categories (id, main, name, icon_name) values (12,false,'Accesorios','');

--Accesorios -- Audio
insert into categories_relations (id_main, id_child) values (10,11);
--Accesorios -- Auriculares
insert into categories_relations (id_main, id_child) values (10,12);

--------------------------------------------------------------------------------------------------------------------------------

--
--PRODUCTS (productos)
--
-- Notebooks id_category = 2
insert into products (id_category,name,description,price,stock,image) values (2,'Notebook lenovo V110','Notebook lenovo V110',10000000,1,'/products/NotebooklenovoV110.png');
insert into products (id_category,name,description,price,stock,image) values (2,'Notebook Dell Inspiron','Notebook Dell Inspiron',10000000,1,'/products/NotebookDellInspiron.png');

-- Impresoras y Accesorios id_category = 3
insert into products (id_category,name,description,price,stock,image) values (3,'Impresora Epson EcoTank L3110','Impresora Epson EcoTank L3110',10000000,1,'/products/EcoTankL3110.png');
insert into products (id_category,name,description,price,stock,image) values (3,'Impresora HP 107W','Impresora HP 107W',10000000,1,'/products/HP107W.png');

-- PC de Escritorio id_category = 4
insert into products (id_category,name,description,price,stock,image) values (4,'A8 Pc Gamer Armada Amd','A8 Pc Gamer Armada Amd',10000000,1,'/products/A8PC.png');
insert into products (id_category,name,description,price,stock,image) values (4,'Pc Gamer Completa Intel Core I3 8100','Pc Gamer Completa Intel Core I3 8100',10000000,1,'/products/I38100.png');

--Smartwatch id_category = 6
insert into products (id_category,name,description,price,stock,image) values (3,'Smartwatch 2.0 LTE Wifi','Smartwatch 2.0 LTE Wifi',10000000,1,'/products/watch-2.png');

--Celulares id_category = 7
insert into products (id_category,name,description,price,stock,image) values (7,'Samsung s20','Ultimo celular de samsung',10000000,1,'/products/samsungs20.png');
insert into products (id_category,name,description,price,stock,image) values (7,'Xiaomi Redmi Note 8T','Xiaomi Redmi Note 8T Dual SIM 64 GB Azul estelar 4 GB RAM',4315000,1,'/products/Note8T.png');
insert into products (id_category,name,description,price,stock,image) values (7,'Moto G6 Plus','Moto G6 Plus 64 GB Nimbus 4 GB RAM',2899900,1,'/products/MotoG6.png');
insert into products (id_category,name,description,price,stock,image) values (7,'iPhone x','iPhone x',10000000,1,'/products/iphone-x.png');
insert into products (id_category,name,description,price,stock,image) values (7,'OnePlus 6T','OnePlus 6T',10000000,1,'/products/oneplus-6t.png');


--Electrodomesticos y aires id_category = 9
insert into products (id_category,name,description,price,stock,image) values (9,'Aire acondicionado LG','Aire acondicionado',4500000,3,'/products/AireLG.png');

-- Audio id_category = 11
insert into products (id_category,name,description,price,stock,image) values (11,'Wireless Audio System','Wireless Audio System',10000000,1,'/products/speakers.png');

-- Auriculares id_category = 12
insert into products (id_category,name,description,price,stock,image) values (12,'Purple Solo 2 Wireless','Purple Solo 2 Wireless',10000000,1,'/products/headphone.png');


--
--OFFERS (ofertas)
--

insert into offers (id_product, discount,expiration,stock) values (1, 30,'20-08-2020',2);
insert into offers (id_product, discount,expiration,stock) values (2, 30,'20-08-2020',2);
insert into offers (id_product, discount,expiration,stock) values (3, 30,'20-08-2020',2);
insert into offers (id_product, discount,expiration,stock) values (4, 30,'20-08-2020',2);
insert into offers (id_product, discount,expiration,stock) values (5, 30,'20-08-2020',2);
insert into offers (id_product, discount,expiration,stock) values (6, 30,'20-08-2020',2);



