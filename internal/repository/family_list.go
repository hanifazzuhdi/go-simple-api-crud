package repository

import (
	"context"
	"simple-api/internal/entity"
)

type FamilyListRepository interface {
	GetAll(ctx context.Context) ([]entity.FamilyList, error)
	GetAllByCustomerID(ctx context.Context, customerID int) ([]entity.FamilyList, error)
	GetByID(ctx context.Context, id int) (entity.FamilyList, error)
	Create(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error)
	Update(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error)
	Delete(ctx context.Context, id int) error
}
