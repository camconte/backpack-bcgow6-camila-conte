package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller (t *testing.T){
	//arrange

	sellerID := "1F"

	productsExpected := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}

	repository := NewRepository()
	
	//act

	productsResult, err := repository.GetAllBySeller(sellerID)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, productsExpected, productsResult)
}