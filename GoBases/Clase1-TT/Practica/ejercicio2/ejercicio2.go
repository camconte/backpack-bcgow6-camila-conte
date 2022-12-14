/*Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000. 
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import "fmt"

func main()  {
	edad := 23
	empleado := true
	anioAntiguedad := 2
	sueldo := 110000

	if edad > 22 && empleado == true && anioAntiguedad > 1 {
		fmt.Println("Se otorga el préstamo")
		if sueldo < 100000 {
			fmt.Println("No se cobrarán intereses")
		}
	} else {
		fmt.Println("No se otorga el péstamo")
	}
}