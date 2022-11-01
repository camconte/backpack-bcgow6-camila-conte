CREATE DATABASE empresa_internet;

CREATE TABLE planes_internet (
id INT PRIMARY KEY NOT NULL,
velocidad_mb INT NOT NULL,
precio FLOAT NOT NULL,
descuento FLOAT);

CREATE TABLE clientes (
dni INT PRIMARY KEY NOT NULL,
nombre VARCHAR(20) NOT NULL,
apellido VARCHAR(20) NOT NULL,
fecha_nacimiento DATE NOT NULL,
provincia VARCHAR(20) NOT NULL,
ciudad VARCHAR(50) NOT NULL,
plan_internet_id INT NOT NULL,
FOREIGN KEY (plan_internet_id) REFERENCES planes_internet(id));

-- Insercion de registros
INSERT INTO planes_internet VALUES (1, 100, 1850, 10);
INSERT INTO planes_internet VALUES (2, 10, 300, 2);
INSERT INTO planes_internet VALUES (3, 50, 1400, 15);
INSERT INTO planes_internet VALUES (4, 25, 1000, 5);
INSERT INTO planes_internet VALUES (5, 5, 200, 0);

INSERT INTO clientes VALUES (40102903, "Juan", "Perez", "2000-01-30", "Salta", "Salta Capital", 2);
INSERT INTO clientes VALUES (44100791, "Martina", "Gonzalez", "2002-05-11", "Formosa", "Formosa", 3);
INSERT INTO clientes VALUES (38701222, "Cristian", "Sanchez", "1996-11-01", "Buenos Aires", "La Plata", 2);
INSERT INTO clientes VALUES (20981234, "Mariela", "Fernandez", "1987-10-12", "Rio Negro", "San Carlos de Bariloche", 5);
INSERT INTO clientes VALUES (12012193, "Sandra", "Perez", "1960-11-21", "Buenos Aires", "CABA", 1);
INSERT INTO clientes VALUES (22093819, "Martin", "Terrena", "1980-07-10", "La Pampa", "Santa Rosa", 4);
INSERT INTO clientes VALUES (30213420, "Santiago", "Ferri", "1998-07-19", "La Pampa", "Macachin", 1);
INSERT INTO clientes VALUES (33434567, "Juan", "Terraza", "1999-08-30", "Buenos Aires", "Bahía Blanca", 5);
INSERT INTO clientes VALUES (40192864, "Lucia", "Gillón", "2001-09-21", "Buenos Aires", "Bolivar", 3);
INSERT INTO clientes VALUES (23430987, "Amanda", "Santenares", "1984-12-10", "Jujuy", "San Salvador de Jujuy", 4);

-- Consultas
-- listar nombre, apellido y fecha de nacimiento de los clientes los cuales su nombre comience con Mar y su año de nacimiento sea mayor a 1985.
SELECT nombre, apellido, fecha_nacimiento FROM clientes
WHERE nombre LIKE "Mar%" AND year(fecha_nacimiento) > 1985;
-- listar nombre, apellido y ciudad de los clientes los cuales su ciudad comience con San y su mes de nacimiento sea noviembre o diciembre.
SELECT nombre, apellido, ciudad FROM clientes
WHERE ciudad LIKE "San%" AND month(fecha_nacimiento) > 10;
-- listar la velocidad en mb y el precio de los planes de internet que tengan un descuento mayor al 7%
SELECT velocidad_mb, precio FROM planes_internet
WHERE descuento > 7;
-- listar el nombre, la provincia y el plan de internet de los clientes que tengan el plan de id 5 o que sean de buenos aires.
SELECT nombre, provincia, plan_internet_id FROM clientes
WHERE plan_internet_id = 5 OR provincia LIKE "Buenos Aires";
-- listar el apellido, provincia y ciudad de los clientes que tengan el plan de internet con id 1.
SELECT apellido, provincia, ciudad FROM clientes
WHERE plan_internet_id = 1;
-- listar todos los campos de los clientes que vivan en buenos aires y ordenarlos por fecha de nacimiento en orden ascendente.
SELECT * FROM clientes
WHERE provincia LIKE "Buenos Aires"
ORDER BY fecha_nacimiento;
-- listar la cantidad de clientes que tienen el plan de internet con id 3
SELECT COUNT(*) cantidad, plan_internet_id FROM clientes
WHERE plan_internet_id = 3;
-- listar la cantidad de clientes que tiene la empresa
SELECT COUNT(*) cantidad FROM clientes;
-- listar el promedio de precios de los planes de internet
SELECT AVG(precio) promedio FROM planes_internet;
-- listar todos los campos de la tabla clientes que residan en bs as y ordenarlos por nombre descendientemente.
SELECT * FROM clientes
WHERE provincia LIKE "Buenos Aires"
ORDER BY nombre DESC;