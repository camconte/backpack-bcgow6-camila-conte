/*La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. 
- Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
- Luego imprimí cada una de las letras.
*/
package main

import "fmt"

func main()  {
	palabra := "Camila"

	fmt.Println(len(palabra))

	for i := 0; i < len(palabra); i++ {
		letra := palabra[i]
		fmt.Printf("%c\n", letra)
	}

	 
}