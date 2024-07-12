CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL,
    is_admin BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS films (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description VARCHAR(1000),
    release_date DATE NOT NULL,
    assesment INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS actors (
    id BIGSERIAL PRIMARY KEY,
    gender VARCHAR(20),
    birthdate DATE,
    name VARCHAR(150)
);

INSERT INTO users (email, encrypted_password, is_admin) VALUES (admin@mail.ru, $2a$04$IpAACgn5/jGuN3ZzZCUJqu727qIC3CtYQYE1iH3BcZdOTX7gvnG.O, true);

CREATE TABLE IF NOT EXISTS films_actors (
    film_id BIGINT NOT NULL,
    actor_id BIGINT NOT NULL,
    PRIMARY KEY (film_id, actor_id),
    FOREIGN KEY (film_id) REFERENCES films (id) ON DELETE CASCADE,
    FOREIGN KEY (actor_id) REFERENCES actors (id) ON DELETE CASCADE
);
