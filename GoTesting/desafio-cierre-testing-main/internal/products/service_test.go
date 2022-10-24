package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySellerSuccess(t *testing.T) {
	//arrange
	sellerID := "1A"

	database := []Product{
		{
			ID: "stub",
			SellerID: "1A",
			Description: "New Product",
			Price: 501.31,
		},
		{
			ID: "mock",
			SellerID: "5H0",
			Description: "limited edition product",
			Price: 300.1,
		},
		{
			ID: "dummy",
			SellerID: "1A",
			Description: "dummy double",
			Price: 100,
		},
	}

	productsExpected := []Product{
		{
			ID: "stub",
			SellerID: "1A",
			Description: "New Product",
			Price: 501.31,
		},
		{
			ID: "dummy",
			SellerID: "1A",
			Description: "dummy double",
			Price: 100,
		},
	}

	mockRepository := MockRepository{
		Data: database,
	} 
	service := NewService(&mockRepository)

	//act

	productsResult, err := service.GetAllBySeller(sellerID)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, productsExpected, productsResult)
}

func TestGetAllBySellerFail(t *testing.T) {
	//arrange
	expectedError := errors.New("no products was found")

	sellerID := "A1F"

	mockRepository := MockRepository{
		Data: nil,
	}
	service := NewService(&mockRepository)

	//act

	productsResult, err := service.GetAllBySeller(sellerID)

	//assert

	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productsResult)
}