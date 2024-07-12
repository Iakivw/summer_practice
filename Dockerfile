FROM postgres:12.2-alpine

COPY ./migrations /docker-entrypoint-initdb.d/
