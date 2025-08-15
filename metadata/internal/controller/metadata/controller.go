package metadata

import (
	"context"
	"errors"

	"github.com/stevan1008/movieGolang/metadata/internal/repository"
	model "github.com/stevan1008/movieGolang/metadata/pkg"
)

// ErrNotFound is returned when a request record is not found.
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a metadata service controller.
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller.
func New(repo metadataRepository) *Controller {
	return &Controller{repo: repo}
}

// Get returns movie metada by id
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrorNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
