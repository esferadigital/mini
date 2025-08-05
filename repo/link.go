package repo

import "github.com/esfxra/mini/entities"

type LinkStorage interface {
	GetLink(key string) (entities.Link, bool)
	CreateAndSaveLink(url string, encodeFunc func(int64) string) (entities.Link, error)
}

type LinkRepository struct {
	storage LinkStorage
}

func NewLinkRepository(storage LinkStorage) *LinkRepository {
	return &LinkRepository{
		storage: storage,
	}
}

func (r *LinkRepository) FindByShortCode(shortCode string) (entities.Link, bool) {
	return r.storage.GetLink(shortCode)
}

func (r *LinkRepository) CreateAndSave(url string, encodeFunc func(int64) string) (entities.Link, error) {
	return r.storage.CreateAndSaveLink(url, encodeFunc)
}
