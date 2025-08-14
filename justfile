run:
    go run cmd/main.go

run-local-database:
    podman run -d --name mini-db -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=adminpass docker.io/library/postgres

migrate-local-database:
    goose -dir infra/migrations postgres://admin:adminpass@localhost:5432/mini_db up

dev-up:
    podman compose -f infra/containers/compose.yaml up -d

dev-down:
    podman compose -f infra/containers/compose.yaml down
