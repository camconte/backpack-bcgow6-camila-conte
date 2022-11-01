package products

import "errors"

type MockRepository struct{
	Data []Product
	
}

func (m *MockRepository) GetAllBySeller(sellerID string) ([]Product, error){
	var productsList []Product

	for _, product := range m.Data {
		if product.SellerID == sellerID {
			productsList = append(productsList, product)
		}
	}

	if len(productsList) == 0{
		return nil, errors.New("no products was found")
	}

	return productsList, nil
}