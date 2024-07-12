package repository

import (
	"context"
	"pr1/internal/rest/model"
)

type BicycleRepository interface {
	Create(ctx context.Context, bicycle model.Bicycle) error
	Read(ctx context.Context, id int64) (model.Bicycle, error)
	Update(ctx context.Context, bicycle model.Bicycle) error
	Delete(ctx context.Context, id int64) error

	List(ctx context.Context) ([]model.Bicycle, error)
}
