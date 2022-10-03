/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
- Usuario: Nombre, Apellido, Correo, Productos (array de productos).
- Producto: Nombre, precio, cantidad.
Se requieren las funciones:
- Nuevo producto: recibe nombre y precio, y retorna un producto.
- Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
- Borrar productos: recibe un usuario, borra los productos del usuario.
*/
package main

import "fmt"

type product struct{
	Name string
	Price float64
	Quantity int
}

type user struct {
	Name string
	Lastname string
	Mail string
	Products []product
}

func newProduct(name string, price float64, quantity int) product{
	return product{
		Name: name,
		Price: price,
		Quantity: quantity,
	}
}

func addProduct(user *user, product product, quantity int) {
	for i := 0; i < quantity; i++ {
		user.Products = append(user.Products, product)
	}
}

func deleteProducts(user *user) {
	user.Products = []product{}
}

func main()  {
	user := user{
		Name: "Camila",
		Lastname: "Conte",
		Mail: "camila@hotmail.com",
	}

	fmt.Printf("%+v\n", user)
	
	product1 := product{
		Name: "Mate",
		Price: 1300,
		Quantity: 2,
	}
	
	product2 := newProduct("Monitor", 50000, 3)
	
	addProduct(&user, product1, 2)
	
	fmt.Printf("%+v\n", user)
	
	addProduct(&user, product2, 1)
	
	fmt.Printf("%+v\n", user)
}