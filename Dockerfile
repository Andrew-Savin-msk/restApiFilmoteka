FROM ubuntu:24.04

# Установка зависимостей
RUN apt-get update && apt-get install -y wget tar curl gnupg lsb-release

# Установка Go
ENV GOLANG_VERSION 1.22.2
RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -O go.tgz && \
    tar -C /usr/local -xzf go.tgz && \
    rm go.tgz

ENV PATH="/usr/local/go/bin:$PATH"

# Установка golang-migrate из бинарного архива
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz -o migrate.tar.gz && \
    tar -xzf migrate.tar.gz -C /usr/local/bin && \
    rm migrate.tar.gz

# Проверка версии migrate
RUN migrate --version

WORKDIR /go/restApiFilmoteka

# Установка переменных окружения для конфигурации
ENV DOCKER_CONFIG_PATH=configs/server_docker.toml

# Копирование исходного кода в контейнер
COPY . .

# Установка зависимостей Go
RUN go mod tidy

# Запуск миграций и основного приложения
CMD migrate -path migrations -database "postgres://postgres:Sassassa12@db:5432/filmoteka?sslmode=disable" up \
  && go run cmd/main.go
