package products

import (
	"fmt"
	"testing"
	"time"
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
		DataMock: initialDatabase,
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

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// Arrange.
	expectedErr := fmt.Errorf("an unexpected error has ocurred")

	mockStorage := MockStorage{
		DataMock:   nil,
		errOnWrite: nil,
		errOnRead:  fmt.Errorf("an unexpected error has ocurred"),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	databaseResult, err := service.GetAll()

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, databaseResult)
}

func TestServiceIntegrationStore(t *testing.T) {
	// Arrange.
	expectedDatabase := []Product{
		{
			Id:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Colour: "Blue",
			Stock: 2000,
			Price: 300,
			Code: "1234",
			Published: true,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
		{
			Id:		2,
			Name:  "Rexona",
			Colour: "Pink",
			Stock: 34,
			Price: 100,
			Code: "456",
			Published: false,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
		{
			Id:    3,
			Name:  "Galletitas Sonrisas 400gr",
			Colour: "White",
			Stock: 600,
			Price: 130,
			Code: "7890",
			Published: true,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
	}

	initialDatabase := []Product{
		{
			Id:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Colour: "Blue",
			Stock: 2000,
			Price: 300,
			Code: "1234",
			Published: true,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
		{
			Id:		2,
			Name:  "Rexona",
			Colour: "Pink",
			Stock: 34,
			Price: 100,
			Code: "456",
			Published: false,
			CreatedAt: time.Now().Format("02-01-2006"),
		},
	}

	mockStorage := MockStorage{
		DataMock: initialDatabase,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	productToCreate := Product{
		Id:    3,
		Name:  "Galletitas Sonrisas 400gr",
		Colour: "White",
		Stock: 600,
		Price: 130,
		Code: "7890",
		Published: true,
		CreatedAt: time.Now().Format("02-01-2006"),
	}

	productCreated, err := service.Store("Galletitas Sonrisas 400gr", "White",130, 600, "7890", true)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, expectedDatabase, mockStorage.DataMock)
	assert.Equal(t, productToCreate, productCreated)
}

func TestServiceIntegrationStoreFailOnLastID(t *testing.T) {
	// Arrange.
	expectedErr := fmt.Errorf("an error has ocurred")

	mockStorage := MockStorage{
		DataMock:   nil,
		errOnRead:  fmt.Errorf("an error has ocurred"),
		errOnWrite: nil,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	/* productToCreate := Product{
		Id:    3,
		Name:  "Galletitas Sonrisas 400gr",
		Colour: "White",
		Stock: 600,
		Price: 130,
		Code: "7890",
		Published: true,
		CreatedAt: time.Now().Format("02-01-2006"),
	} */

	productCreated, err := service.Store("Galletitas Sonrisas 400gr", "White",130, 600, "7890", true)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Empty(t, productCreated)
}

func TestServiceIntegrationStoreFailOnRepository(t *testing.T) {
	// Arrange.
	expectedErr := fmt.Errorf("an error has ocurred")

	//Case 1 - errOnRead
	mockStorageReadError := MockStorage{
		DataMock:   nil,
		errOnRead:  fmt.Errorf("an error has ocurred"),
		errOnWrite: nil,
	}

	repository1 := NewRepository(&mockStorageReadError)
	service1 := NewService(repository1)

	//Case 2 - errOnWrite
	mockStorageWriteError := MockStorage{
		DataMock:   nil,
		errOnRead:  nil,
		errOnWrite: fmt.Errorf("an error has ocurred"),
	}

	repository2 := NewRepository(&mockStorageWriteError)
	service2 := NewService(repository2)

	// Act.

	//Case 1 - errOnRead
	productCreated1, err1 := service1.Store("Galletitas Sonrisas 400gr", "White",130, 600, "7890", true)

	//Case 2 - errOnWrite
	productCreated2, err2 := service2.Store("Galletitas Sonrisas 400gr", "White",130, 600, "7890", true)
	
	// Assert.
	
	//Case 1 - errOnRead
	assert.EqualError(t, err1, expectedErr.Error())
	assert.Empty(t, productCreated1)

	//Case 2 - errOnWrite
	assert.EqualError(t, err2, expectedErr.Error())
	assert.Empty(t, productCreated2)
}


func TestServiceIntegrationUpdateNameAndPrice(t *testing.T) {
	//arrange
	productToUpdate := Product{
		Id: 1,
		Name: "Coffee",
		Price: 340,
	}

	//product expected
	productUpdated := Product{
		Id: 1,
		Name: "New Coffee",
		Price: 120,
	}

	productsStorage := []Product{productToUpdate}

	myMockStore := MockStorage{
		DataMock: productsStorage,
	}

	repository := NewRepository(&myMockStore)
	service := NewService(repository)

	//act

	productResult, err := service.UpdateNameAndPrice(productToUpdate.Id, "New Coffee", 120)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, productUpdated, productResult)
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
		DataMock: initialDatabase,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	
	//act

	//case 1 - errOnRead - id exists
	err0 := service.Delete(2)
	databaseResult, err1 := service.GetAll()

	//case 2 - errOnWrite - id not exists
	err2 := service.Delete(3)

	//assert
	
	//case 1 - errOnRead
	assert.Nil(t, err0)
	assert.Nil(t, err1)
	assert.Equal(t, expectedDatabase, databaseResult)

	//case 2 - errOnWrite
	assert.EqualError(t, err2, expectedErr.Error())
}

