package service

import (
	"context"
	"simple-api/internal/entity"
	"simple-api/internal/repository"
)

type customerServiceImpl struct {
	customerRepository   repository.CustomerRepository
	familyListRepository repository.FamilyListRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository, familyListRepository repository.FamilyListRepository) CustomerService {
	return &customerServiceImpl{
		customerRepository:   customerRepository,
		familyListRepository: familyListRepository,
	}
}

func (service customerServiceImpl) GetAll(ctx context.Context) ([]entity.Customer, error) {
	customers, err := service.customerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (service customerServiceImpl) GetById(ctx context.Context, id int) (entity.Customer, error) {
	customer, err := service.customerRepository.GetById(ctx, id)
	if err != nil {
		return entity.Customer{}, err
	}

	familyLists, err := service.familyListRepository.GetAllByCustomerID(ctx, customer.CstId)
	if err != nil {
		return entity.Customer{}, err
	}

	customer.FamilyList = familyLists

	return customer, nil
}

func (service customerServiceImpl) Create(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	createdCustomer, err := service.customerRepository.Create(ctx, customer)
	if err != nil {
		return customer, err
	}

	return createdCustomer, nil
}

func (service customerServiceImpl) Update(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	updatedCustomer, err := service.customerRepository.Update(ctx, customer)
	if err != nil {
		return customer, err
	}

	return updatedCustomer, nil
}

func (service customerServiceImpl) Delete(ctx context.Context, id int) error {
	return service.customerRepository.Delete(ctx, id)
}

func (service customerServiceImpl) SyncCustomerFamilies(ctx context.Context, customerId int, familyList []entity.FamilyList) (entity.Customer, error) {
	customer, err := service.customerRepository.GetById(ctx, customerId)
	if err != nil {
		return entity.Customer{}, err
	}

	for _, family := range familyList {
		if family.FlId <= 0 {
			_, errCreate := service.familyListRepository.Create(ctx, entity.FamilyList{
				FlName:     family.FlName,
				FlDob:      family.FlDob,
				FlRelation: family.FlRelation,
				CstId:      customerId,
			})

			if errCreate != nil {
				return entity.Customer{}, err
			}
		} else {
			_, errUpdate := service.familyListRepository.Update(ctx, entity.FamilyList{
				FlId:       family.FlId,
				FlName:     family.FlName,
				FlDob:      family.FlDob,
				FlRelation: family.FlRelation,
				CstId:      customerId,
			})

			if errUpdate != nil {
				return entity.Customer{}, err
			}
		}
	}

	return customer, nil
}
