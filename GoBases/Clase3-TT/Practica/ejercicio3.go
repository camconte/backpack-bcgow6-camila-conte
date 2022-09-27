/*Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
- Productos: nombre, precio, cantidad.
- Servicios: nombre, precio, minutos trabajados.
- Mantenimiento: nombre, precio.

Se requieren 3 funciones:
- Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
- Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
- Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/

package main

import "fmt"

type product struct{
	Name string
	Price float64
	Quantity int
}

type service struct {
	Name string
	Price float64
	WorkMinutes float64
}

type maintenance struct{
	Name string
	Price float64
}

func sumProducts (ch chan float64, products []product) (total float64){
	for _, product := range products {
		total += (product.Price * float64(product.Quantity))
	}
	ch <- total
	return
}

func sumServices (ch chan float64, services []service) (total float64){
	for _, service := range services {
		if service.WorkMinutes < 30 {
			total += service.Price
		}else if service.WorkMinutes >= 30 {
			total += service.Price * (service.WorkMinutes/30)
		}
	}
	ch <- total
	return
}

func sumMaintenance(ch chan float64, maintenances []maintenance) (total float64) {
	for _, maintenance := range maintenances {
		total += maintenance.Price
	}
	ch <- total
	return
}

func main()  {
	product1 := product{
		Name: "prod1",
		Price: 20,
		Quantity: 2,
	}

	product2 := product{
		Name: "prod2",
		Price: 10,
		Quantity: 3,
	}

	service1 := service{
		Name: "serv1",
		Price: 10,
		WorkMinutes: 120,
	}

	service2 := service{
		Name: "serv2",
		Price: 30,
		WorkMinutes: 28,
	}

	maintenance1 := maintenance{
		Name: "maintenance",
		Price: 45,
	}

	sliceProducts := []product{product1, product2}
	sliceServices:= []service{service1, service2}
	sliceMaintenance := []maintenance{maintenance1}
	
	//productChannel := make(chan product)
	//serviceChannel := make(chan service)
	//maintenanceChannel := make(chan maintenance)

	totalChannel := make(chan float64)

	var totalSales float64

	go sumProducts(totalChannel, sliceProducts)
	totalProducts := <- totalChannel
	
	go sumServices(totalChannel, sliceServices)
	totalServices := <- totalChannel

	go sumMaintenance(totalChannel, sliceMaintenance)
	totalMaintenances := <- totalChannel

	fmt.Println("Total Products : ", totalProducts)
	fmt.Println("Total Services : ", totalServices)
	fmt.Println("Total Maintenances : ", totalMaintenances)
	
	totalSales = totalMaintenances + totalProducts + totalServices

	fmt.Println("Total Sales : ", totalSales,". \nProgram finished")

}
