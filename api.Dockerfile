FROM golang:1.14-alpine AS build

ARG APP_NAME
ARG APP_VER

ENV GO111MODULE=auto
ENV APP_NAME=$APP_NAME
ENV APP_VER=$APP_VER
ENV APP_PORT=80
ENV APP_MODE=release
ENV COCKCROACH_DB_CONN_STR=postgresql://root@localhost:26257/postgress?sslmode=disable

WORKDIR /app/api
COPY ./api/bin/. .

ENTRYPOINT ["./api"]

EXPOSE $APP_PORT 