-- Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
SELECT d.nombre_depto, e.puesto, d.localidad FROM departamentos d
INNER JOIN empleados e ON e.depto_nro = d.depto_nro;

-- Visualizar los departamentos con más de cinco empleados.
SELECT COUNT(*) total_empleados, d.nombre_depto FROM departamentos d
INNER JOIN empleados e ON e.depto_nro = d.depto_nro
GROUP BY d.depto_nro
HAVING total_empleados > 5;

-- Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el 
-- mismo puesto que ‘Mito Barchuk’.
SELECT d.nombre_depto, e.nombre, e.salario FROM departamentos d
INNER JOIN empleados e ON e.depto_nro = d.depto_nro
WHERE e.puesto LIKE "Presidente";

-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
SELECT * FROM empleados e
INNER JOIN departamentos d ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto LIKE "Contabilidad"
ORDER BY e.nombre;

-- Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT salario, nombre FROM empleados
ORDER BY salario ASC
LIMIT 1;


-- Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e.nombre, e.apellido, e.salario, d.nombre_depto FROM empleados e
INNER JOIN departamentos d ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto LIKE "Ventas"
ORDER BY e.salario DESC
LIMIT 1;

SELECT MAX(e.salario) salarioMax, d.nombre_depto, CONCAT(e.nombre, " ", e.apellido) empleado FROM empleados e
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
GROUP BY d.nombre_depto, empleado
HAVING d.nombre_depto LIKE "Ventas"
ORDER BY salarioMax DESC
LIMIT 1;

-- Insercion de datos
INSERT INTO departamentos VALUES ("D-000-1", "Software", "Los Tigres");
INSERT INTO departamentos VALUES ("D-000-2", "Sistemas", "Guadalupe");
INSERT INTO departamentos VALUES ("D-000-3", "Contabilidad", "La Roca");
INSERT INTO departamentos VALUES ("D-000-4", "Ventas", "Plata");

INSERT INTO empleados VALUES ("E-0001", "Cesar", "Piñero", "Vendedor", "2018-05-12", 80000, 15000, "D-000-4");
INSERT INTO empleados VALUES ("E-0002", "Yosep", "Kowaleski", "Analista", "2015-07-14", 140000, 0, "D-000-2");
INSERT INTO empleados VALUES ("E-0003", "Mariela", "Barrios", "Director", "2014-06-05", 185000, 0, "D-000-3");
INSERT INTO empleados VALUES ("E-0004", "Jonathan", "Aguileta", "Vendedor", "2015-06-03", 85000, 10000, "D-000-4");
INSERT INTO empleados VALUES ("E-0005", "Daniel", "Brezezicki", "Vendedor", "2018-03-03", 83000, 10000, "D-000-4");
INSERT INTO empleados VALUES ("E-0006", "Mito", "Barchuk", "Presidente", "2014-06-05", 190000, 0, "D-000-3");
INSERT INTO empleados VALUES ("E-0007", "Emilio", "Galarza", "Desarrollador", "2014-08-02", 60000, 0, "D-000-1");