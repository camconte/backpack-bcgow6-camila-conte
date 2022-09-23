/*Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares.
- Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
- Imprime los valores de las variables en consola.
- ¿Qué tipo de dato le asignarías a las variables?
*/

package main

import "fmt"

var temperatura int = 17
var humedad int = 44
var presion float64 = 1021.0

func main() {
	fmt.Println("La temperatura actual es ", temperatura,"º", " con una humedad del ", humedad, "% y una presión de ", presion, "mb")
}
