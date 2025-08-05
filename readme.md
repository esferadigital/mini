# mini
Reference implementation of a very basic URL shortener service with Go.

## Brief
The URL shortener assigns each new URL a sequential ID number (1, 2, 3...), then encodes that number into a short code. This approach ensures every short code is unique since the counter only moves forward.

**Flow:**
1. New URL arrives → Counter increments (e.g., from 42 to 43)
2. ID gets encoded → Produces short code (e.g., "aB3x")
3. Short code maps to original URL in storage

This mimics how database auto-increment IDs work, but keeps everything in memory for simplicity.

## Characteristics
- Domain driven design
- Dependency injection
- In-memory storage

## Test the program
```bash
# Run the program
go run main.go

# Create a short URL:
curl -X POST http://localhost:8080/shorten -d '{ "url": "https://www.theverge.com/openai/718785/openai-gpt-oss-open-model-release" }'
```

Then visit the short URL in your browser.
