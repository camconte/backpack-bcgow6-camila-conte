package products

import (
	"encoding/json"
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

	productsResult, _ := repository.GetAll()

	//assert
	assert.Equal(t, productsExpected, productsResult)
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

	productResult, _ := repository.UpdateNameAndPrice(productToUpdate.Id, "New Coffee", 120)

	//assert

	assert.Equal(t, productUpdated, productResult)
	assert.True(t, myMockStore.ReadMethodWasCall)
}




