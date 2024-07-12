FROM ubuntu:24.04

RUN apt-get update && apt-get install -y wget tar curl gnupg lsb-release

ENV GOLANG_VERSION 1.22.2
RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -O go.tgz && \
    tar -C /usr/local -xzf go.tgz && \
    rm go.tgz

ENV PATH="/usr/local/go/bin:$PATH"

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz -o migrate.tar.gz && \
    tar -xzf migrate.tar.gz -C /usr/local/bin && \
    rm migrate.tar.gz

RUN migrate --version

WORKDIR /go/restApiFilmoteka

ENV DOCKER_CONFIG_PATH=configs/server_docker.toml

COPY . .

RUN go mod tidy

CMD migrate -path migrations -database "postgres://postgres:Sassassa12@db:5432/filmoteka?sslmode=disable" up \
  && go run cmd/main.go
