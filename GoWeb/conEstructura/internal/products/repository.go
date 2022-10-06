package products

import (
	"fmt"
	"github.com/camconte/backpack-bcgow6-camila-conte/GoWeb/conEstructura/pkg/store"
)

type Product struct{
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Colour string `json:"colour" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Stock int `json:"stock" binding:"required"`
	Code string `json:"code" binding:"required"`
	Published bool `json:"published" binding:"required"`
	CreatedAt string `json:"createdAt"`
}

var productsStorage []Product
//var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, colour string, price float64, stock int, code string, published bool, createdAt string) (Product, error)
	LastID() (int, error)
	Update(id int, name string, colour string, price float64, stock int, code string, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct{
	db store.Store
} //implementa los metodos de la interfaz

//devuelve el repo
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error){
	var productsArray []Product

	err := r.db.Read(&productsArray)
	if err != nil{
		return nil, err
	}
	
	return productsArray, nil
}

func (r *repository) Store(id int, name string, colour string, price float64, stock int, code string, published bool, createdAt string) (Product, error){
	var productsArray []Product

	err := r.db.Read(&productsArray)
	if err != nil{
		return Product{}, err
	}

	p := Product{id, name, colour, price, stock, code, published, createdAt}


	productsArray = append(productsArray, p)

	if err := r.db.Write(productsArray); err != nil{
		return Product{}, err
	}
	
	return p, nil
}

func (r *repository) LastID() (int, error){
	var productsArray []Product

	err := r.db.Read(&productsArray)
	if err != nil{
		return 0, err
	}

	if len(productsArray) == 0 {
		return 0, nil
	}

	return productsArray[len(productsArray)-1].Id, nil
}

//faltan que actualicen el archivo los siguientes metodos:
func (r *repository) Update(id int, name string, colour string, price float64, stock int, code string, published bool) (Product, error){
	p := Product{
		Name: name,
		Colour: colour,
		Price: price,
		Stock: stock,
		Code: code,
		Published: published,
	}

	//chequeamos si existe para actualizar el valor correspondiente
	updated := false
	for i, product := range productsStorage {
		if product.Id == id{
			p.Id = id
			p.CreatedAt = product.CreatedAt
			productsStorage[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("product with id %d not found", id)
	}

	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error){
	var updatedProduct Product

	updated := false

	for _, product := range productsStorage {
		if product.Id == id{
			if name != "" {
				product.Name = name
				updated = true
			}
			if price != 0 {
				product.Price = price
				updated = true
			}

			updatedProduct = product
		}
	}

	if !updated {
		return updatedProduct, fmt.Errorf("product with id %d not found", id)
	}

	return updatedProduct, nil
}

func (r *repository) Delete(id int) error {
	deleted := false

	for i, product := range productsStorage {
		if product.Id == id{
			//lo eliminamos del storage
			productsStorage = append(productsStorage[:i], productsStorage[i+1:]...)
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("product with id %d not found", id)
	}

	return nil
}