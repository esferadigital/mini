package services

import (
	"context"
	"log"

	"github.com/esferachill/mini/internal/platform"
)

func Target(slug string) (string, error) {
	ctx := context.Background()
	platform := platform.GetPlatform()
	shortURL, err := platform.DBClient.Queries.GetShortURLBySlug(ctx, slug)
	if err != nil {
		log.Printf("error getting short url by slug: %v\n", err)
		return "", err
	}
	return shortURL.TargetUrl, nil
}
