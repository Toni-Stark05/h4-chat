#!/bin/bash

until pg_isready -h db -p 5432 -U user; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 2
done

goose -dir /app/migrations postgres "postgres://user:password@db:5432/mydb?sslmode=disable" up

/goapp
