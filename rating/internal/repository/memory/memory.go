package memory

import (
	"context"
	"errors"

	model "github.com/stevan1008/movieGolang/rating/pkg"
)

var ErrorNotFound = errors.New("error not found")

// Repository defines a rating repository.
type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

// New creates anew memory repository.
func New() *Repository {
	return &Repository{map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

// Get retrieves all ratings for a given record.
func (r *Repository) Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, ErrorNotFound
	}
	if records, ok := r.data[recordType][recordID]; !ok || len(records) == 0 {
		return nil, ErrorNotFound
	}
	return r.data[recordType][recordID], nil
}

func (r *Repository) Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}
	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)
	return nil
}
