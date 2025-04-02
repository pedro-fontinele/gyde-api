create table product(
	id serial primary key,
	name varchar(150) not null,
	price numeric(18, 2) not null
);

select *
from product;

INSERT INTO product (name, price) VALUES
('Smartphone Samsung Galaxy S23', 4999.99),
('Notebook Dell Inspiron 15', 3799.90),
('Smart TV LG 55" OLED', 6999.99),
('Console PlayStation 5', 4299.00),
('Fone de Ouvido Bluetooth JBL', 299.99),
('Mouse Gamer Logitech G502', 349.90),
('Teclado Mecânico Razer BlackWidow', 699.99),
('Monitor Gamer 27" Acer', 1899.90),
('Cadeira Gamer DXRacer', 1599.00),
('Processador AMD Ryzen 9 5900X', 2399.90),
('Placa de Vídeo RTX 4070 Ti', 4999.99),
('HD Externo 2TB Seagate', 549.90),
('SSD NVMe 1TB Kingston', 799.99),
('Impressora Multifuncional HP', 899.00),
('Câmera DSLR Canon EOS Rebel T7', 2499.00),
('Smartwatch Apple Watch Series 8', 4299.90),
('Tablet Samsung Galaxy Tab S8', 3899.90),
('Echo Dot 5ª Geração', 379.99),
('Ventilador de Mesa Arno', 199.90),
('Cafeteira Expresso Nespresso', 499.90);
