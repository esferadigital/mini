package services

import (
	"context"
	"time"

	"github.com/esferachill/mini/internal/platform"
	"github.com/esferachill/mini/internal/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

func RecordVisit(slug string, userAgent string, occurredAt time.Time) (repo.Visit, error) {
	db := platform.GetPlatform().DBClient
	visit, err := db.Queries.CreateVisit(context.Background(), repo.CreateVisitParams{
		Slug:       slug,
		UserAgent:  pgtype.Text{String: userAgent, Valid: true},
		OccurredAt: pgtype.Timestamptz{Time: occurredAt, Valid: true},
	})
	if err != nil {
		return repo.Visit{}, err
	}
	return visit, nil
}
