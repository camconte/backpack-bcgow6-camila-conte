package mocks

import (
	"fmt"

	"github.com/camconte/backpack-bcgow6-camila-conte/GoTesting/Clase2-TM/internal/products"
)

type MockRepository struct{
	DataMock []products.Product
	Error error
}

func (m *MockRepository) Update(id int, name string, colour string, price float64, stock int, code string, published bool) (products.Product, error){
	p := products.Product{
		Name: name,
		Colour: colour,
		Price: price,
		Stock: stock,
		Code: code,
		Published: published,
	}
	
	updated := false
	for i, product := range m.DataMock {
		if product.Id == id{
			p.Id = id
			p.CreatedAt = product.CreatedAt
			m.DataMock[i] = p
			updated = true
		}
	}

	if !updated {
		return products.Product{}, fmt.Errorf("product with id %d not found", id)
	}

	return p, nil
}