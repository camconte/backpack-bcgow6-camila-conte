/*
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/
package main

import (
	"fmt"
)


func throwSalaryError(salary int) (err error) {
	if salary < 150_000 {
		err = fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de %d", salary)
	}else{
		fmt.Println("Debe pagar impuesto")
	}
	return 
}

func main()  {
	
	salary := 140_000

	err := throwSalaryError(salary)

	if err != nil {
		fmt.Println(err)
	}

}