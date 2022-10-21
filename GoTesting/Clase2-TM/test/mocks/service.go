package mocks

import "github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/internal/products"


type MockService struct{
	repository MockRepository
}

func (m *MockService) Update(id int, name string, colour string, price float64, stock int, code string, published bool) (products.Product, error){
	return m.repository.Update(id, name, colour, price, stock, code, published)
}
