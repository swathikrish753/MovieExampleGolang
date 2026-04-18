package repository

import (
	"context"
	"sync"

	model "movieexample.com/metadata/pkg"
)

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

	if metadata, ok := r.data[id]; ok {
		return metadata, nil
	}
	return nil, ErrNotFound
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
