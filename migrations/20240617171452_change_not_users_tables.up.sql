ALTER TABLE actors DROP CONSTRAINT IF EXISTS film_id;
ALTER TABLE actors DROP COLUMN IF EXISTS film_id;

CREATE TABLE IF NOT EXISTS films_actors (
    film_id BIGINT NOT NULL,
    actor_id BIGINT NOT NULL,
    PRIMARY KEY (film_id, actor_id),
    FOREIGN KEY (film_id) REFERENCES films (id) ON DELETE CASCADE,
    FOREIGN KEY (actor_id) REFERENCES actors (id) ON DELETE CASCADE
);
