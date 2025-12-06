package repository

import (
	"sync"

	model "movieexample.com/metadata/pkg"
)

// Repository defines a memory movie metadata repository.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func New() *Repository {
	return &Repository{
		data: make(map[string]*model.Metadata),
	}
}

func (r *Repository) Get(id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, ErrNotFound
	}
	return m, nil
}

func (r *Repository) Put(id string, m *model.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = m
	return nil
}
