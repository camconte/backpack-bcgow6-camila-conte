package products

import (
	"database/sql"
	"fmt"

	"clase1/internal/domain"
)

type Repository interface {
	Store(p domain.Product) (int, error)
	GetByName(name string) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetProductsByWarehouse(warehouseId int) ([]domain.Product, error)
	Update(product domain.Product, id int) error 
	Delete(id int) error 
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

const (
	STORE_PRODUCT = "INSERT INTO products (name, type, count, price, id_warehouse) VALUES (?,?,?,?,?)"

	GET_BY_NAME = "SELECT id, name, type, count, price FROM products WHERE name = ?;"

	GET_ALL = "SELECT id, name, type, count, price FROM products"

	DELETE_PRODUCT = "DELETE FROM products WHERE id = ?"

	UPDATE_PRODUCT = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?"

	GET_ALL_PRODUCTS_BY_WAREHOUSE = "SELECT p.id, p.name, p.type, p.count, p.price FROM products p INNER JOIN warehouses w ON p.id_warehouse = w.id WHERE w.id = ?"
)

func (r *repository) GetProductsByWarehouse(warehouseId int) ([]domain.Product, error){
	rows, err := r.db.Query(GET_ALL_PRODUCTS_BY_WAREHOUSE, warehouseId)
	if err != nil {
		return []domain.Product{}, fmt.Errorf("error al ejecutar la consulta - %v", err)
	}

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.ProductType, &product.Count, &product.Price); err != nil {
			return []domain.Product{}, fmt.Errorf("no se encontraron registros - error %v", err)
		}

		products = append(products, product)
	}

	return products, nil
}

/* Ejercicio 2 - Replicar Store()
Tomar el ejemplo visto en la clase y diseñar el método Store():
Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
Implementar el método Store.
*/
func (r *repository) Store(p domain.Product) (int, error) {
	stmt, err := r.db.Prepare(STORE_PRODUCT)
	if err != nil {
		return 0, fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.ProductType, p.Count, p.Price, p.WarehouseId)
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener último id - error %v", err)
	}

	return int(id), nil
}

/* Ejercicio 1 - Implementar GetByName()
Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:
Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorna una estructura del tipo Product.
Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.
*/
func (r *repository) GetByName(name string) (domain.Product, error) {

	row := r.db.QueryRow(GET_BY_NAME, name)
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.ProductType, &product.Count, &product.Price); err != nil {
		return domain.Product{}, fmt.Errorf("no registros para %s - error %v", name, err)
	}

	return product, nil
}


/*Ejercicio 1- Clase2 TM
Diseñar un método GetAll.
Dentro del archivo repository desarrollar el método GetAll().
Comprobar el correcto funcionamiento.
*/
func (r *repository) GetAll() ([]domain.Product, error){
	var products []domain.Product

	rows, err := r.db.Query(GET_ALL);
	if err != nil {
		return []domain.Product{}, fmt.Errorf("error: ha ocurrido un problema en la consulta - %v", err)
	}

	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(&product.Id, &product.Name, &product.ProductType, &product.Count, &product.Price); err != nil {
			return []domain.Product{}, fmt.Errorf("no se encontraron registros - error %v", err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Update(product domain.Product, id int) error {
	stm, err := r.db.Prepare(UPDATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stm.Close()


	result, err := stm.Exec(product.Name, product.ProductType, product.Count, product.Price, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected < 1 {
		return fmt.Errorf("error: no affected rows")
	}

	return nil
}


/*Ejercicio 2 - Clase 2 TM 
Diseñar un método para eliminar un recurso de la base de datos.
Dentro del archivo repository desarrollar el método Delete().
Comprobar el correcto funcionamiento.
*/
func (r *repository) Delete(id int) error{
	stmt, err := r.db.Prepare(DELETE_PRODUCT)
	if err != nil {
		return fmt.Errorf("error al preparar la consulta - error %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta - error %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las lineas afectadas - error %v", err)
	}

	if rowsAffected < 1 {
		return fmt.Errorf("error al eliminar el producto")
	}

	return nil
}