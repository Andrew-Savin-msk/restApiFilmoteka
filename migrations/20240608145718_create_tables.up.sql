CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL
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
    gender VARCHAR(20) NOT NULL,
    birthdate DATE NOT NULL,
    film_id INTEGER,
    FOREIGN KEY(film_id) REFERENCES films(id)
);