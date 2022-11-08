package product

import (
	"clase1/internal/domain"

	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

//el test deberia hacerse sobre un mock y no sobre una bd real
func TestStore(t *testing.T){
	//arrange
	product := domain.Product{
		Name: "Lays Papas",
		ProductType: "Snacks",
		Count: 34,
		Price: 120.4,
	}

	repo := NewRepo()

	//act
	lastID, err := repo.Store(product)
	if err != nil {
		log.Println(err)
	}

	//assert
	assert.Equal(t, int64(8), lastID)
}

func TestGetByName(t *testing.T){
	//arrange
	searchName := "aceitunas"
	product := domain.Product{
		Id: 5,
		Name: "aceitunas",
		ProductType: "snacks",
		Count: 12,
		Price: 30.6,
	}

	repo := NewRepo()

	//act
	productResult, err := repo.GetByName(searchName)
	if err != nil {
		log.Println(err)
	}

	//assert
	assert.Equal(t, product, productResult)

}