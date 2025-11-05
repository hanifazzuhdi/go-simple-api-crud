package service

import (
	"context"
	"simple-api/internal/entity"
)

type CustomerService interface {
	GetAll(ctx context.Context) ([]entity.Customer, error)
	GetById(ctx context.Context, id int) (entity.Customer, error)
	Create(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	Update(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	Delete(ctx context.Context, id int) error
	SyncCustomerFamilies(ctx context.Context, customerId int, familyList []entity.FamilyList) (entity.Customer, error)
}
