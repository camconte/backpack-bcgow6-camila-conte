package products

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

/* ------------------------------- Ejercicio 1 ------------------------------ */
type StubStore struct{
	//simula ser el archivo
	Products []Product
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

func (s StubStore) Read(data interface{}) error {

	productsJSON, _ := json.Marshal(s.Products)

	//almaceno lo que esta en el array de products en la data que me llega por parametro
	return json.Unmarshal(productsJSON, &data)
}

func TestGetAll(t *testing.T){
	//arrange
	var productsExpected []Product
	
	product1 := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}
	product2 := Product{
		Id: 2,
		Name: "Mobile",
		Colour: "Grey",
		Price: 1100,
		Stock: 4,
		Code: "34566",
	}

	productsExpected = append(productsExpected, product1)
	productsExpected = append(productsExpected, product2)

	myStubStore := StubStore{productsExpected}
	repository := NewRepository(myStubStore)

	//Act

	productsResult, err := repository.GetAll()

	//assert
	assert.Nil(t, err)
	assert.Equal(t, productsExpected, productsResult)
}

func TestGetAllFail(t *testing.T) {
	//arrange
	var database []Product
	
	product1 := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}
	product2 := Product{
		Id: 2,
		Name: "Mobile",
		Colour: "Grey",
		Price: 1100,
		Stock: 4,
		Code: "34566",
	}

	database = append(database, product1)
	database = append(database, product2)

	expectedError := fmt.Errorf("an unexpected error has ocurred")
	
	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: fmt.Errorf("an unexpected error has ocurred"),
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)
	
	//act

	productsResult, err := repository.GetAll()

	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, productsResult)
}
/* ------------------------------- Ejercicio 2 ------------------------------ */
type MockStore struct{
	BeforeUpdate Product
	ReadMethodWasCall bool
	AfterUpdate Product
	Products []Product
}

func (m *MockStore) Write(data interface{}) error {
	//en lugar de usar el package json se podria "castear" con punteros. Ejemplo: data.(*[]Product)
	/*dataProducts := data.(*[]Product)
	m.Products = *dataProducts*/
	dataJSON, _ := json.Marshal(data)
	return json.Unmarshal(dataJSON, &m.Products)
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadMethodWasCall = true
	productsJSON, _ := json.Marshal(m.Products)

	return json.Unmarshal(productsJSON, &data)
}

func TestUpdateNameAndPrice(t *testing.T) {
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

	myMockStore := MockStore{
		BeforeUpdate: productToUpdate,
		ReadMethodWasCall: false,
		AfterUpdate: productUpdated,
		Products: productsStorage,
	}

	repository := NewRepository(&myMockStore)

	//act

	productResult, err := repository.UpdateNameAndPrice(productToUpdate.Id, "New Coffee", 120)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, productUpdated, productResult)
	assert.True(t, myMockStore.ReadMethodWasCall)
}

/* ------------------------- Clase3 - code coverage ------------------------- */
func TestStore(t *testing.T) {
	//arrange
	var database []Product

	product1 := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}
	product2 := Product{
		Id: 2,
		Name: "Mobile",
		Colour: "Grey",
		Price: 1100,
		Stock: 4,
		Code: "34566",
	}
	product3 := Product{
		Id: 3,
		Name: "IPad",
		Colour: "White",
		Price: 1800,
		Stock: 1,
		Code: "0066",
		CreatedAt: "13:30",
	}

	database = append(database, product1)
	database = append(database, product2)
	database = append(database, product3)

	initialDatabase := []Product{product1, product2}

	myMockStorage := MockStorage{
		DataMock: initialDatabase,
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	productCreated, err := repository.Store(3, "IPad", "White", 1800,1, "0066", false, "13:30")
	
	//assert
	assert.Nil(t, err)
	assert.Equal(t, database, myMockStorage.DataMock)
	assert.Equal(t, product3, productCreated)

}

func TestStoreFailReading(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred while reading")

	myMockStorage := MockStorage{
		DataMock: nil,
		errOnRead: fmt.Errorf("an unexpected error has ocurred while reading"),
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	productCreated, err := repository.Store(3, "IPad", "White", 1800,1, "0066", false, "13:30")
	
	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productCreated)

}

func TestStoreFailWriting(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred while writing")

	myMockStorage := MockStorage{
		DataMock: nil,
		errOnRead: nil,
		errOnWrite: fmt.Errorf("an unexpected error has ocurred while writing"),
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	productCreated, err := repository.Store(3, "IPad", "White", 1800,1, "0066", false, "13:30")
	
	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productCreated)
}

func TestLastID(t *testing.T) {
	//arrange
	var database []Product

	product1 := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}
	product2 := Product{
		Id: 2,
		Name: "Mobile",
		Colour: "Grey",
		Price: 1100,
		Stock: 4,
		Code: "34566",
	}
	product3 := Product{
		Id: 3,
		Name: "IPad",
		Colour: "White",
		Price: 1800,
		Stock: 1,
		Code: "0066",
		CreatedAt: "13:30",
	}

	database = append(database, product1)
	database = append(database, product2)
	database = append(database, product3)

	myMockStorage := MockStorage{
		DataMock: database,
	}

	expectedID := 3

	repository := NewRepository(&myMockStorage)

	//act
	
	lastIDResult, err := repository.LastID()

	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedID, lastIDResult)
}

func TestLastIDFailReading(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred")

	var database []Product

	product1 := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}

	database = append(database, product1)

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: fmt.Errorf("an unexpected error has ocurred"),
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)

	//act
	
	lastIDResult, err := repository.LastID()

	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, lastIDResult)
}

func TestLastIDWithEmptyArray(t *testing.T) {
	//arrange
	var database []Product

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: nil,
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)

	//act
	
	lastIDResult, err := repository.LastID()

	//assert
	assert.Nil(t, err)
	assert.Empty(t, lastIDResult)
}

func TestUpdateFailReading(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred")

	var database []Product

	productToUpdate := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}

	database = append(database, productToUpdate)

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: fmt.Errorf("an unexpected error has ocurred"),
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)

	/* productUpdated := Product{
		Id: 1,
		Name: "New Coffee",
		Colour: "Black",
		Price: 180,
		Stock: 4,
		Code: "1234",
	} */

	//act
	
	productResult, err := repository.Update(1, "New Coffee", "Black", 180, 4, "1234", true)

	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productResult)
}

func TestUpdateFailWithIncorrectID(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("product with id 2 not found")

	var database []Product

	productToUpdate := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}

	database = append(database, productToUpdate)

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: nil,
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)

	/* productUpdated := Product{
		Id: 1,
		Name: "New Coffee",
		Colour: "Black",
		Price: 180,
		Stock: 4,
		Code: "1234",
	} */

	//act
	
	productResult, err := repository.Update(2, "New Coffee", "Black", 180, 4, "1234", true)

	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productResult)
}

func TestUpdateFailWriting(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred")

	var database []Product

	productToUpdate := Product{
		Id: 1,
		Name: "Coffee",
		Colour: "Black",
		Price: 120,
		Stock: 4,
		Code: "1234",
	}

	database = append(database, productToUpdate)

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: nil,
		errOnWrite: fmt.Errorf("an unexpected error has ocurred"),
	}

	repository := NewRepository(&myMockStorage)

	/* productUpdated := Product{
		Id: 1,
		Name: "New Coffee",
		Colour: "Black",
		Price: 180,
		Stock: 4,
		Code: "1234",
	} */

	//act
	
	productResult, err := repository.Update(1, "New Coffee", "Black", 180, 4, "1234", true)

	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productResult)
}

func TestUpdateNameAndPriceFailReading(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred while reading")

	database := []Product{
		{
			Id: 1,
			Name: "pendrive",
			Colour: "grey",
		},
	}

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: fmt.Errorf("an unexpected error has ocurred while reading"),
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	productUpdated, err := repository.UpdateNameAndPrice(1, "IPad", 1800)
	
	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productUpdated)
}

func TestUpdateNameAndPriceFailWithIncorrectID(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("product with id 2 not found")

	database := []Product{
		{
			Id: 1,
			Name: "pendrive",
			Colour: "grey",
		},
	}

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: nil,
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	productUpdated, err := repository.UpdateNameAndPrice(2, "IPad", 1800)
	
	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productUpdated)
}

func TestUpdateNameAndPriceFailWriting(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred while writing")

	database := []Product{
		{
			Id: 1,
			Name: "pendrive",
			Colour: "grey",
		},
	}

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: nil,
		errOnWrite: fmt.Errorf("an unexpected error has ocurred while writing"),
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	productUpdated, err := repository.UpdateNameAndPrice(1, "IPad", 1800)
	
	//assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, productUpdated)
}

func TestDeleteFailReading(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred while reading")

	database := []Product{
		{
			Id: 1,
			Name: "pendrive",
			Colour: "grey",
		},
	}

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: fmt.Errorf("an unexpected error has ocurred while reading"),
		errOnWrite: nil,
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	err := repository.Delete(1)
	
	//assert
	assert.EqualError(t, err, expectedError.Error())
}

func TestDeleteFailWriting(t *testing.T) {
	//arrange
	expectedError := fmt.Errorf("an unexpected error has ocurred while writing")

	database := []Product{
		{
			Id: 1,
			Name: "pendrive",
			Colour: "grey",
		},
	}

	myMockStorage := MockStorage{
		DataMock: database,
		errOnRead: nil,
		errOnWrite: fmt.Errorf("an unexpected error has ocurred while writing"),
	}

	repository := NewRepository(&myMockStorage)
	
	//act
	err := repository.Delete(1)
	
	//assert
	assert.EqualError(t, err, expectedError.Error())
}