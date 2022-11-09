package products

import "clase1/internal/domain"

type Service interface {
	Store(domain.Product) (int, error)
	GetByName(name string) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetProductsByWarehouse(warehouseId int) ([]domain.Product, error)
	Update(domain.Product, int) error
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetProductsByWarehouse(warehouseId int) ([]domain.Product, error){
	return s.repository.GetProductsByWarehouse(warehouseId)
}

func (s *service) Store(p domain.Product) (int, error) {
	return s.repository.Store(p)
}

func (s *service) GetByName(name string) (domain.Product, error) {
	return s.repository.GetByName(name)
}

func (s *service) GetAll() ([]domain.Product, error){
	return s.repository.GetAll()
}

func (s *service) Update(product domain.Product, id int) error {
	return s.repository.Update(product, id)
}

func (s *service) Delete(id int) error{
	return s.repository.Delete(id)
}
