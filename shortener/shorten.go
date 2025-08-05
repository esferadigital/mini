package shortener

import (
	"github.com/esfxra/mini/repo"
)

func Shorten(linkRepo *repo.LinkRepository, url string) string {
	link, _ := linkRepo.CreateAndSave(url, EncodeBase62)
	return link.ShortCode
}

func GetOriginalURL(linkRepo *repo.LinkRepository, shortCode string) (string, bool) {
	link, exists := linkRepo.FindByShortCode(shortCode)
	if !exists {
		return "", false
	}
	return link.URL, true
}
