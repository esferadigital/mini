#!/bin/bash
set -euo pipefail

# Wait for database to be ready (basic retry logic)
echo "Running migrations..."
retry_count=0
max_retries=30

while [ $retry_count -lt $max_retries ]; do
    if goose -dir infra/migrations postgres "$DATABASE_URL" up; then
        echo "Migrations completed successfully"
        break
    else
        echo "Migration attempt $((retry_count + 1)) failed, retrying in 2 seconds..."
        sleep 2
        retry_count=$((retry_count + 1))
    fi
done

if [ $retry_count -eq $max_retries ]; then
    echo "Failed to run migrations after $max_retries attempts"
    exit 1
fi

echo "Starting application..."
exec ./main
