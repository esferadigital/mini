# mini
Reference implementation of a basic URL shortener service with Go.

**Characteristics**

- Domain driven design
- Dependency injection
- Database storage with Postgres, pgx and sqlc

**Short URL generation**

The `crypto/rand` package is used to generate a cryptographically secure random number, which is then encoded into a short code. For instance, a short URL with the current length of `6` looks like `XCE7koad`.

## Tools
```bash
# goose - for migrations
go install github.com/pressly/goose/v3/cmd/goose@latest

# sqlc - to generate models and queries
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# podman - for containerization
sudo pacman -S podman podman-compose

# just - for running tasks
sudo pacman -S just
```

> Make sure `$GOPATH` is configured and added to your PATH.

> These instructions are written for Arch Linux. Check your distribution's package manager for the right tools.

## Test the program
```bash
# Run the program with a local database container
just dev-up

# Listen to logs
podman logs mini-api -f

# Create a short URL:
curl -X POST http://localhost:8080/shorten -d '{ "url": "https://www.theverge.com/openai/718785/openai-gpt-oss-open-model-release" }'
```

Then visit the short URL returned by the program in your browser.

> You could also run the program with a database you already have the credentials for with `just run`
