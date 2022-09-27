/*
Una empresa que se encarga de vender productos de limpieza necesita:
- Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados 	por punto y coma (csv).
- Debe tener el id del producto, precio y la cantidad.
- Estos valores pueden ser hardcodeados o escritos en duro en una variable.

Un archivo CSV es cualquier archivo de texto en el cual los caracteres están separados por comas, haciendo una especie de tabla en filas y columnas. Las columnas quedan definidas por cada punto y coma (;), mientras que cada fila se define mediante una línea adicional en el texto.
*/
package main

import (
	"fmt"
	"os"
)

type product struct {
	ID int
	Price float64
	Quantity int
}

func main()  {
	//el archivo debe tener el id, el precio y la cantidad
	product1 := product{
		ID: 1,
		Price: 1300.15,
		Quantity: 2,
	}

	product2 := product{
		ID: 2,
		Price: 230.1,
		Quantity: 1,
	}


	file := fmt.Sprintf("%d;%.3f;%d\n%d;%.2f;%d\n", product1.ID, product1.Price, product1.Quantity, product2.ID, product2.Price, product2.Quantity)
	

	err := os.WriteFile("./products.csv", []byte(file), 0644)

	if err != nil{
		panic(err)
	}else{
		fmt.Println("File Written successfully!")
	}




}