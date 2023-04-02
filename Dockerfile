FROM postgres:15

ENV POSTGRES_PASSWORD 1234

ENV POSTGRES_USER my_user

ENV POSTGRES_DB my_db

ENV SCHEMA_NAME my_schema

RUN apt update &&  apt install -y postgis && rm -rf /var/lib/apt/lists/*

COPY init_db.sh /docker-entrypoint-initdb.d/

VOLUME ["/var/run/postgresql/"]

EXPOSE 5430