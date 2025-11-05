package repository

import (
	"context"
	"simple-api/internal/entity"
)

type NationalityRepository interface {
	GetAll(ctx context.Context) ([]entity.Nationality, error)
	GetById(ctx context.Context, id int) (entity.Nationality, error)
	Create(ctx context.Context, nationality entity.Nationality) (entity.Nationality, error)
	Update(ctx context.Context, nationality entity.Nationality) (entity.Nationality, error)
	Delete(ctx context.Context, id int) error
}
