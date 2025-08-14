package services

import (
	"context"
	"log"

	"github.com/esferachill/mini/internal/platform"
	"github.com/esferachill/mini/internal/repo"
)

const (
	SlugLength = 6
)

func Shorten(url string) string {
	ctx := context.Background()
	platform := platform.GetPlatform()
	slug, err := GenerateRandomBytes(SlugLength)
	if err != nil {
		// Handle error
		log.Printf("error generating random bytes: %v\n", err)
		return ""
	}

	encodedSlug := EncodeBase62(slug)

	shortURL, err := platform.DBClient.Queries.CreateShortURL(ctx, repo.CreateShortURLParams{
		Slug:      encodedSlug,
		TargetUrl: url,
	})
	if err != nil {
		// Handle error
		log.Printf("error creating short URL: %v\n", err)
		return ""
	}
	return shortURL.Slug
}
