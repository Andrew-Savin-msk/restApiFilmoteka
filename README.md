Конечно! Вот примерная структура README файла для проекта фильмотеки:

---

# Film Library API

A RESTful API for managing a movie database, built with Go and PostgreSQL. The API allows users to manage actors and films, as well as perform searches and retrieve sorted lists of films.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Features

- User authentication and session management
- CRUD operations for actors and films
- Search films by name fragment
- Retrieve sorted lists of films
- Retrieve actors with their films

## Installation

### Prerequisites

- Go 1.22.1 or later
- PostgreSQL 13 or later
- Docker (optional, for containerized setup)

### Clone the Repository

```bash
git clone https://github.com/Andrew-Savin-msk/rest-api-filmoteka.git
cd rest-api-filmoteka
```

### Build and Run

#### Using Docker

1. Build and run the containers using Docker Compose:

    ```bash
    docker-compose up -d
    ```

2. The API will be available at `http://localhost:8081`.

#### Manually

1. Install dependencies:

    ```bash
    go mod tidy
    ```

2. Set up the PostgreSQL database:

    ```bash
    createdb filmoteka
    psql -d filmoteka -c "CREATE USER postgres WITH PASSWORD 'your_password';"
    ```

3. Run the migrations:

    ```bash
    migrate -path migrations_docker -database "postgres://postgres:your_password@localhost:5432/filmoteka?sslmode=disable" up
    ```

4. Run the API server:

    ```bash
    go run cmd/main.go
    ```

15. The API will be available at `http://localhost:8081`.

## Usage

### Configuration

Configuration is managed via TOML files. The default configuration files are located in the `configs` directory.

### Environment Variables

- `CONFIG_PATH`: Path to the main configuration file (default: `configs/server.toml`).
- `CONFIG_PATH_DOCKER`: Path to the Docker-specific configuration file (default: `configs/server_docker.toml`).

## API Endpoints

### User Endpoints

- **Create User**
    - `POST /register`
    - Request Body: `{ "email": "user@example.com", "password": "password" }`
    - Response: Created user object

- **Get Session (Login)**
    - `POST /authorize`
    - Request Body: `{ "email": "user@example.com", "password": "password" }`
    - Response: Status OK

- **Who Am I**
    - `GET /private/who-am-i`
    - Response: Authenticated user object

### Actor Endpoints

- **Create Actor**
    - `POST /private/create-actor`
    - Request Body: `{ "name": "Actor Name", "gender": "M", "birthdate": "01-02-2000" }`
    - Response: Created actor ID

- **Get Actor**
    - `GET /get-actor`
    - Request Body: `{ "id": 1 }`
    - Response: Actor object

- **Delete Actor**
    - `DELETE /private/delete-actor`
    - Request Body: `{ "id": 1 }`
    - Response: Deleted actor ID

- **Update Actor**
    - `PUT/PATCH /private/update-actor`
    - Request Body: `{ "id": 1, "name": "Updated Name", "gender": "F", "birthdate": "01-02-2000" }`
    - Response: Status OK

- **Get All Actors**
    - `GET /get-actors`
    - Response: List of actors with their films

### Film Endpoints

- **Create Film**
    - `POST /private/post-film`
    - Request Body: `{ "name": "Film Name", "description": "Description", "release_date": "01-02-2000", "assesment": 8.5, "actors": [1, 2, 3] }`
    - Response: Created film ID

- **Delete Film**
    - `DELETE /private/delete-film`
    - Request Body: `{ "id": 1 }`
    - Response: Deleted film ID

- **Update Film**
    - `PUT /private/update-film`
    - Request Body: `{ "id":1, "name": "Updated Name", "description": "Updated Description", "release_date": "01-02-2000", "assesment": 9.0 }`
    - Response: Status OK

- **Find Film by Name Part**
    - `PATCH /films`
    - Request Body: `{ "name_part": "Film" }`
    - Response: List of matching films

- **Get Sorted Films**
    - `GET /select-films`
    - Request Body: `{ "sorting_parameter": "name" }`
    - Response: List of films sorted by criteria
