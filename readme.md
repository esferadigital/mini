# mini
Reference implementation of a basic URL shortener service with Go.

**Characteristics**

- Domain driven design
- Dependency injection
- Database storage with Postgres, pgx and sqlc

**Short URL generation**

The URL shortener assigns each new URL a sequential ID number (1, 2, 3...), then encodes that number into a short code. This approach ensures every short code is unique since the counter only moves forward.

**Flow**

1. New URL arrives → Counter increments (e.g., from 42 to 43)
2. ID gets encoded → Produces short code (e.g., "aB3x")
3. Short code maps to original URL in storage

This mimics how database auto-increment IDs work, but keeps everything in memory for simplicity.

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

# Run the program with a database you already have the credentials for
just run

# Create a short URL:
curl -X POST http://localhost:8080/shorten -d '{ "url": "https://www.theverge.com/openai/718785/openai-gpt-oss-open-model-release" }'
```

Then visit the short URL returned by the program in your browser.
