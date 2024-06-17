DROP TABLE IF EXISTS films_actors;

ALTER TABLE actors ADD COLUMN film_id BIGINT;
ALTER TABLE actors ADD CONSTRAINT actors_film_id_fkey FOREIGN KEY (film_id) REFERENCES films (id);
