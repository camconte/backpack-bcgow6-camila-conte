package product

import (
	"clase1/internal/domain"
	"clase1/pkg/db"
)

type Repository interface {
	GetByName(name string) (domain.Product, error)
	Store(product domain.Product) (int64, error)
}

//no recibe nada ya que utiliza la variable StorageDB directamente desde el package (no es lo recomendable)
type repository struct {}

func NewRepo() Repository {
	return &repository{}
}

const (
	GET_NAME_PRODUCT = "SELECT * FROM products WHERE name = ?"
	STORE_PRODUCT = "INSERT INTO products(name, type, count, price) VALUES (?, ?, ?, ?)"
)

func (r *repository) GetByName(name string)(product domain.Product, err error){
	db := db.StorageDB

	rows := db.QueryRow(GET_NAME_PRODUCT, name)
	if err := rows.Scan(&product.Id, &product.Name, &product.ProductType, &product.Count, &product.Price);err != nil {
		return domain.Product{}, err
	}
	return 
}

func (r *repository) Store(product domain.Product) (lastID int64, err error){
	db := db.StorageDB

	statement, err := db.Prepare(STORE_PRODUCT)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(product.Name, product.ProductType, product.Count, product.Price)
	if err != nil {
		return 0, err
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return
}


