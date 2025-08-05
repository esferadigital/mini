package infra

import (
	"sync"

	"github.com/esfxra/mini/entities"
)

type Memory struct {
	mu           sync.RWMutex
	links        map[string]entities.Link
	linksCounter int64
}

func NewMemory() *Memory {
	return &Memory{
		links:        make(map[string]entities.Link),
		linksCounter: 0,
	}
}

func (m *Memory) GetLink(key string) (entities.Link, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	link, exists := m.links[key]
	return link, exists
}

// CreateAndSaveLink atomically increments the counter, creates a link, and saves it
func (m *Memory) CreateAndSaveLink(url string, encodeFunc func(int64) string) (entities.Link, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.linksCounter++
	shortCode := encodeFunc(m.linksCounter)

	link := entities.Link{
		ID:        m.linksCounter,
		URL:       url,
		ShortCode: shortCode,
	}

	m.links[shortCode] = link
	return link, nil
}
