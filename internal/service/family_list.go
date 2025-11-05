package service

import (
	"context"
	"simple-api/internal/entity"
)

type FamilyListService interface {
	GetAll(ctx context.Context) ([]entity.FamilyList, error)
	GetById(ctx context.Context, id int) (entity.FamilyList, error)
	Create(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error)
	Update(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error)
	Delete(ctx context.Context, id int) error
	GetFamilyMemberByCustomerID(ctx context.Context, customerID int) ([]entity.FamilyList, error)
}
