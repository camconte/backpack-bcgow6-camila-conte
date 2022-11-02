-- PRIMERA PARTE
-- 5. Escribir una consulta genérica para cada uno de los siguientes diagramas:
/*SELECT * FROM leftTable
INNER JOIN rightTable ON leftTable.id = rightTable.id*/
/*SELECT * FROM leftTable
LEFT JOIN rightTable ON leftTable.id = rightTable.id*/

-- SEGUNDA PARTE
-- Mostrar el título y el nombre del género de todas las series.
SELECT s.title, g.name genre FROM series s
INNER JOIN genres g ON g.id = s.genre_id;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT e.title, a.first_name, a.last_name FROM episodes e
INNER JOIN actor_episode ae ON e.id = ae.episode_id
INNER JOIN actors a ON a.id = ae.actor_id;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT s.title, COUNT(*) total_seasons FROM series s
INNER JOIN seasons seas ON s.id = seas.serie_id
GROUP BY s.title;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, 
-- siempre que sea mayor o igual a 3.
SELECT g.name, COUNT(*) total_movies FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.name
HAVING total_movies >= 3;

-- Mostrar sólo el nombre y apellido de los actores que trabajan 
-- en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT a.first_name, a.last_name FROM actors a
INNER JOIN actor_movie am ON am.actor_id = a.id
INNER JOIN movies m ON m.id = am.movie_id
WHERE m.title LIKE "La Guerra de las galaxias%";

