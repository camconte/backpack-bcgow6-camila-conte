-- Listar los datos de los autores.
SELECT * FROM autores;
-- Listar nombre y edad de los estudiantes
SELECT nombre, edad FROM estudiantes;
-- ¿Qué estudiantes pertenecen a la carrera informática?
SELECT * FROM estudiantes
WHERE carrera LIKE "%informática%";
-- ¿Qué autores son de nacionalidad francesa o italiana?
SELECT * FROM autores
WHERE nacionalidad LIKE "Francesa" OR nacionalidad LIKE "Italiana";
-- ¿Qué libros no son del área de internet?
SELECT * FROM libros
WHERE area NOT LIKE "%Internet%";
-- Listar los libros de la editorial Salamandra.
SELECT * FROM libros
WHERE editorial LIKE "%Salamandra%";
-- Listar los datos de los estudiantes cuya edad es mayor al promedio.

-- pendiente
SELECT AVG(edad) promedio, estudiantes.edad FROM estudiantes
GROUP BY id_lector
HAVING estudiantes.edad > promedio;

-- Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
SELECT nombre, apellido FROM estudiantes
WHERE apellido LIKE "G%";
-- Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
SELECT a.nombre FROM autores a
INNER JOIN libro_autor la ON la.id_autor = a.id_autor
INNER JOIN libros l ON l.id_libro = la.id_libro
WHERE l.titulo LIKE "El Universo: Guía de viaje";
-- ¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT * FROM libros l
INNER JOIN prestamos p ON p.libros_id_libro = l.id_libro
INNER JOIN estudiantes e ON p.estudiantes_id_lector = e.id_lector
WHERE e.nombre LIKE "Filippo" AND e.apellido LIKE "Galli";
-- Listar el nombre del estudiante de menor edad.
SELECT edad, nombre FROM estudiantes
ORDER BY edad ASC
LIMIT 1;
-- Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT e.nombre FROM estudiantes e
INNER JOIN prestamos p ON p.estudiantes_id_lector = e.id_lector
INNER JOIN libros l on p.libros_id_libro = l.id_libro
WHERE l.area LIKE "Base de Datos";
-- Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT l.*, a.nombre FROM libros l
INNER JOIN libro_autor la ON la.id_libro = l.id_libro
INNER JOIN autores a ON la.id_autor = a.id_autor
WHERE a.nombre LIKE "J.K. Rowling"; 
-- Listar títulos de los libros que debían devolverse el 16/07/2021.
SELECT l.titulo FROM libros l
INNER JOIN prestamos p ON p.libros_id_libro = l.id_libro
WHERE p.fecha_devolucion = "2021-07-16";


-- Insercion de datos
INSERT INTO autores VALUES (1, "J.K. Rowling", "Británica");
INSERT INTO autores VALUES (2, "Ruperto Zcheta", "Francesa");
INSERT INTO autores VALUES (3, "Paulo Serra", "Italiana");
INSERT INTO autores VALUES (4, "Michelle Ricci", "Italiana");
INSERT INTO autores VALUES (5, "Nuevo autor", "Argentina");

INSERT INTO estudiantes VALUES (1, "Juan", "Perez", "direccion 123", "Ingeniería informática", 17);
INSERT INTO estudiantes VALUES (2, "Fernando", "Flores", "direccion 123", "Ingeniería electrónica", 25);
INSERT INTO estudiantes VALUES (3, "Juana", "Gervasio", "direccion 123", "Contabilidad", 20);
INSERT INTO estudiantes VALUES (4, "Tiago", "Solari", "direccion 123", "Ingeniería informática", 28);
INSERT INTO estudiantes VALUES (5, "Mabel", "Serrano", "direccion 123", "Educación física", 22);
INSERT INTO estudiantes VALUES (6, "Filippo", "Galli", "direccion 123", "Abogacía", 21);

INSERT INTO libros VALUES (1, "Harry Potter", "Las nuevas vibras", "Entretenimiento");
INSERT INTO libros VALUES (2, "SQL vs NoSQL", "Salamandra", "Base de Datos");
INSERT INTO libros VALUES (3, "Nuevo mundo", "Salamandra", "Internet");
INSERT INTO libros VALUES (4, "El Universo: Guía de viaje", "Las nuevas vibras", "Entretenimiento");
INSERT INTO libros VALUES (5, "Otro Libro", "Nueva editorial", "Base de Datos");

INSERT INTO libro_autor VALUES (1, 1);
INSERT INTO libro_autor VALUES (2, 4);
INSERT INTO libro_autor VALUES (3, 5);
INSERT INTO libro_autor VALUES (1, 4);
INSERT INTO libro_autor VALUES (4, 3);

INSERT INTO prestamos VALUES (6, 2, "2021-07-10", "2021-07-16", 0);
INSERT INTO prestamos VALUES (3, 1, "2021-07-20", "2021-07-26", 1);
INSERT INTO prestamos VALUES (4, 5, "2021-08-09", "2021-08-25", 1);
INSERT INTO prestamos VALUES (2, 3, "2021-07-10", "2021-07-16", 0);
INSERT INTO prestamos VALUES (1, 4, "2021-10-10", "2021-10-16", 1);



