/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
  - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
  - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/
package main

import "fmt"

const(
	SMALL = "Small"
	MEDIUM = "Medium"
	BIG = "Big"
)


type ecommerce struct{
	products []Product
}

func (e ecommerce) Total() float64 {
	var total float64

	for _, product := range e.products {
		total += product.CostCalculator()
	}

	return total
}

func (e *ecommerce) Add(p product){
	e.products = append(e.products, &p)
}

type product struct{
	ProductType string
	Name string
	Price float64
}

func (p *product) CostCalculator() float64 {
	cost := p.Price
	switch p.ProductType {
		case MEDIUM:
			cost = p.Price + (p.Price*0.3)
		case BIG:
			cost = p.Price + (p.Price*0.6) + 2500
	}

	return cost
}


type Product interface {
	CostCalculator() float64
}

type Ecommerce interface {
	Total() float64
	Add(product)
}

func newProduct(productType string, name string, price float64) product {
	return product{
		Name: name,
		Price: price,
		ProductType: productType,
	}
}

func newEcommerce() Ecommerce{
	return &ecommerce{}
}

func main()  {
	ecommerce := newEcommerce()
	product1 := newProduct(MEDIUM, "MacBook", 1200)
	product2 := newProduct(BIG, "Heladera", 30000)
	product3 := newProduct(SMALL, "Mate", 300)

	ecommerce.Add(product1)
	ecommerce.Add(product2)
	ecommerce.Add(product3)
	
	fmt.Println(ecommerce.Total())

}
