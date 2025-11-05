package service

import (
	"context"
	"simple-api/internal/entity"
	"simple-api/internal/repository"
)

type familyListServiceImpl struct {
	familyListRepository repository.FamilyListRepository
}

func NewFamilyListService(familyListRepository repository.FamilyListRepository) FamilyListService {
	return &familyListServiceImpl{
		familyListRepository: familyListRepository,
	}
}

func (service familyListServiceImpl) GetAll(ctx context.Context) ([]entity.FamilyList, error) {
	familyLists, err := service.familyListRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return familyLists, nil
}

func (service familyListServiceImpl) GetById(ctx context.Context, id int) (entity.FamilyList, error) {
	familyList, err := service.familyListRepository.GetByID(ctx, id)
	if err != nil {
		return entity.FamilyList{}, err
	}

	return familyList, nil
}

func (service familyListServiceImpl) Create(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error) {
	createdFamilyList, err := service.familyListRepository.Create(ctx, familyList)
	if err != nil {
		return familyList, err
	}

	return createdFamilyList, nil
}

func (service familyListServiceImpl) Update(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error) {
	updatedFamilyList, err := service.familyListRepository.Update(ctx, familyList)
	if err != nil {
		return familyList, err
	}

	return updatedFamilyList, nil
}

func (service familyListServiceImpl) Delete(ctx context.Context, id int) error {
	return service.familyListRepository.Delete(ctx, id)
}

func (service familyListServiceImpl) GetFamilyMemberByCustomerID(ctx context.Context, customerID int) ([]entity.FamilyList, error) {
	familyLists, err := service.familyListRepository.GetAllByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return familyLists, nil
}
