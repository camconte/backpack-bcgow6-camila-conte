package products

import "time"

type Service interface{
	GetAll() ([]Product, error)
	Store(name string, colour string, price float64, stock int, code string, published bool) (Product, error)
	Update(id int, name string, colour string, price float64, stock int, code string, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type service struct{
	repository Repository
}

//devuelve el service
func NewService(r Repository) Service{
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error){
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(name string, colour string, price float64, stock int, code string, published bool) (Product, error){
	lastID, err := s.repository.LastID()
	if err != nil{
		return Product{}, err
	}

	lastID++

	createdAt := time.Now().Format("02-01-2006")

	product, err := s.repository.Store(lastID, name, colour, price, stock, code, published, createdAt)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, name string, colour string, price float64, stock int, code string, published bool) (Product, error){
	return s.repository.Update(id, name, colour, price, stock, code, published)
}

func (s *service) UpdateNameAndPrice(id int, name string, price float64) (Product, error){
	return s.repository.UpdateNameAndPrice(id, name, price)
}

func (s *service) Delete(id int) error{
	return s.repository.Delete(id)
}