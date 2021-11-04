# syntax=docker/dockerfile:1

FROM postgres:latest

WORKDIR /

COPY *.sql /docker-entrypoint-initdb.d/

EXPOSE 8080