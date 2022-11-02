-- Agregar una película a la tabla movies.
INSERT INTO movies (title, rating, awards, release_date, length, genre_id) VALUES ("Avengers: Endgame", 7.3, 4, "2019-04-26", 181, 5);

-- Agregar un género a la tabla genres.
INSERT INTO genres (name, ranking, active) VALUES ("Anime", 13, 1);

-- Asociar a la película del punto 1. con el género creado en el punto 2.
UPDATE movies SET genre_id = 16 
WHERE title="Avengers: Endgame";

-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.
UPDATE actors SET favorite_movie_id = 22 
WHERE id = 2;

-- Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE copy_movies 
SELECT * FROM movies;

-- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM copy_movies WHERE awards < 5;

SELECT * FROM copy_movies;

-- Obtener la lista de todos los géneros que tengan al menos una película.
SELECT COUNT(*) total_movies, g.name FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.id
HAVING total_movies >= 1;

-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT a.*, m.title, m.awards FROM actors a
INNER JOIN movies m ON m.id = a.favorite_movie_id
WHERE m.awards > 3;

-- Crear un índice sobre el nombre en la tabla movies.
ALTER TABLE movies ADD INDEX idx_movie_title (title);

-- Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;

-- En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.
/*si ya que el campo title es un dato muy utilizado para las consultas debido a que es el campo mediante el cual se identifica cada pelicula*/

-- ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta
/*quiza podria mejorar la perfomance de las consultas un índice en la columna name de la tabla genre ya que puede llegar a utilizarse para relacionar más fácilmente cada genero con la pelicula que lo tenga*/