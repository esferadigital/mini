default:
    just --list

run:
    go run cmd/main.go

run-container:
    podman run -d --name mini-api -p 8080:8080 -e "DATABASE_URL=postgres://admin:adminpass@localhost:5432/mini_db" mini-api

run-database:
    podman run -d --name mini-db -p 5432:5432 -e POSTGRES_PASSWORD=adminpass -e POSTGRES_USER=admin -e POSTGRES_DB=mini_db postgres:17

build:
    go build -o tmp/main cmd/main.go

build-container:
    podman build -t mini-api -f infra/containers/local/Containerfile .

migrate-local-database:
    goose -dir infra/migrations postgres://admin:adminpass@localhost:5432/mini_db up

dev-up:
    podman compose -f infra/containers/local/compose.yaml up -d

dev-down:
    podman compose -f infra/containers/local/compose.yaml down

generate-sql:
    sqlc generate

generate-components:
    templ generate
