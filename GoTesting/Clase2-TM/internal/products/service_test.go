package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	//arrange
	product1 := Product{
		Id: 1,
		Name: "Rick Coffee",
		Price: 200,
	}

	product2 := Product{
		Id: 2,
		Name: "Strange Lollipop",
		Price: 80,
	}

	initialDatabase := []Product{product1, product2}

	product2Updated := Product{
		Id: 2,
		Name: "Strange Cola",
		Price: 180,
		Colour: "Purple",
		Code: "123A",
		Stock: 20,
	}

	//expectedDatabase := []Product{product1,product2Updated}


	mockStorage := MockStorage{
		dataMock: initialDatabase,
		ReadMethodWasCall: false,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	//act
	productResult, err := service.Update(2,"Strange Cola","Purple", 180, 20, "123A", false)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, product2Updated, productResult)
	assert.True(t, mockStorage.ReadMethodWasCall)
	
}

func TestServiceIntegrationDelete(t *testing.T) {
	//arrange
	product1 := Product{
		Id: 1,
		Name: "Rick Coffee",
		Price: 200,
	}

	product2 := Product{
		Id: 2,
		Name: "Strange Lollipop",
		Price: 80,
	}

	initialDatabase := []Product{product1, product2}

	expectedDatabase := []Product{product1}

	expectedErr := fmt.Errorf("product with id 3 not found")

	mockStorage := MockStorage{
		dataMock: initialDatabase,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	
	//act

	//case 1 - id exists
	err0 := service.Delete(2)
	databaseResult, err1 := service.GetAll()

	//case 2 - id not exists
	err2 := service.Delete(3)

	//assert
	
	//case 1
	assert.Nil(t, err0)
	assert.Nil(t, err1)
	assert.Equal(t, expectedDatabase, databaseResult)

	//case 2
	assert.EqualError(t, err2, expectedErr.Error())
}

