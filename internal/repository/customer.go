package repository

import (
	"context"
	"simple-api/internal/entity"
)

type CustomerRepository interface {
	GetAll(ctx context.Context) ([]entity.Customer, error)
	GetById(ctx context.Context, id int) (entity.Customer, error)
	Create(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	Update(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	Delete(ctx context.Context, id int) error
}
