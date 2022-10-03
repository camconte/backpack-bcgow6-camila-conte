/*Realizar una aplicación que contenga una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
Ej: 7, Julio

Se podria hacer con varios if/else if pero elijo el switch ya que es más recomendado su uso en casos como el actual
*/

package main

import "fmt"

func main()  {
	nroMes := 3

	switch nroMes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")	
	case 5:
		fmt.Println("Mayo")	
	case 6:
		fmt.Println("Junio")	
	case 7:
		fmt.Println("Julio")	
	case 8:
		fmt.Println("Agosto")	
	case 9:
		fmt.Println("Septiembre")	
	case 10:
		fmt.Println("Octubre")	
	case 11:
		fmt.Println("Noviembre")	
	case 12:
		fmt.Println("Diciembre")	
	}
}