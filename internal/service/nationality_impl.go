package service

import (
	"context"
	"simple-api/internal/entity"
	"simple-api/internal/repository"
)

type nationalityServiceImpl struct {
	nationalityRepository repository.NationalityRepository
}

func NewNationalityService(nationalityRepository repository.NationalityRepository) NationalityService {
	return &nationalityServiceImpl{
		nationalityRepository: nationalityRepository,
	}
}

func (service nationalityServiceImpl) GetAll(ctx context.Context) ([]entity.Nationality, error) {
	nationalities, err := service.nationalityRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return nationalities, nil
}

func (service nationalityServiceImpl) GetById(ctx context.Context, id int) (entity.Nationality, error) {
	nationality, err := service.nationalityRepository.GetById(ctx, id)
	if err != nil {
		return entity.Nationality{}, err
	}

	return nationality, nil
}

func (service nationalityServiceImpl) Create(ctx context.Context, nationality entity.Nationality) (*entity.Nationality, error) {
	nationality, err := service.nationalityRepository.Create(ctx, nationality)
	if err != nil {
		return nil, err
	}

	return &nationality, nil
}

func (service nationalityServiceImpl) Update(ctx context.Context, nationality entity.Nationality) (*entity.Nationality, error) {
	nationality, err := service.nationalityRepository.Update(ctx, nationality)
	if err != nil {
		return nil, err
	}

	return &nationality, nil
}

func (service nationalityServiceImpl) Delete(ctx context.Context, id int) error {
	err := service.nationalityRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
