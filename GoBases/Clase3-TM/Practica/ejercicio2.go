/*La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	data, err := os.ReadFile("./products.csv")
	if err != nil{
		panic(err)
	}
	fmt.Printf("ID\t\t    Precio\tCantidad\n")
	
	lines := strings.Fields(string(data))

	var total float64
	for _, line := range lines {
		product := strings.Split(line, ";")
		fmt.Printf("%s\t\t%10s\t%7s\n", product[0], product[1], product[2])
		price, err := strconv.ParseFloat(product[1], 32)
		quantity, err := strconv.ParseFloat(product[2], 32)
		if err != nil {
			panic(err)
		}
		total += price * quantity
	}

	fmt.Printf("\t\t%10.3f\n", total)




	/*firstComa := true
	//var total float64
	for _, char := range string(data) {
		if string(char) == ";" && firstComa{
			fmt.Printf("\t\t")
			firstComa = false
		}else if string(char) == ";" && !firstComa{
			fmt.Printf("\t")
			firstComa = true
		}else{
			fmt.Print(string(char))
		}
	}*/

	

	//fmt.Println(string(data))
}