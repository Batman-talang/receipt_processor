package storage

import (
	"github.com/google/uuid"
	"sync"
)

type MemoryStore struct {
	mu      sync.RWMutex
	records map[string]int
}

func New() *MemoryStore {
	return &MemoryStore{
		records: make(map[string]int),
	}
}

func (s *MemoryStore) Save(points int) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.NewString()
	s.records[id] = points
	return id
}
func (s *MemoryStore) Get(id string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	point, ok := s.records[id]
	return point, ok
}

func (s *MemoryStore) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, _ := range s.records {
		delete(s.records, k)
	}
}
