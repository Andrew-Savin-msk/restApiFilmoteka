CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    encrypted_password VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS films (
    id INTEGER PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description VARCHAR(1000),
    assesment INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS actors (
    id INTEGER PRIMARY KEY,
    gender VARCHAR(20) NOT NULL,
    birthday DATE NOT NULL,
    film_id INTEGER,
    FOREIGN KEY(film_id) REFERENCES films(id)
);