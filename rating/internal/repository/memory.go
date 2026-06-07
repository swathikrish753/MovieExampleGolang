package memory

import (
	"context"
	"errors"

	model "movieexample.com/rating/pkg"
)

type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

func New() *Repository {
	return &Repository{
		map[model.RecordType]map[model.RecordID][]model.Rating{},
	}

}

func (r *Repository) Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, errors.New("Record not found")
	}
	if ratings, ok := r.data[recordType][recordID]; !ok || len(ratings) == 0 {
		return nil, errors.New("Record not found")
	}
	return r.data[recordType][recordID], nil
}

func (r *Repository) Put(ctx context.Context, recordType model.RecordType, recordID model.RecordID, recordVal model.RatingValue, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}
	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)
	return nil
}
