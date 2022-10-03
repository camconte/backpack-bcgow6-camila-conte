/*
1. En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/
package main

import (
	"fmt"
	"os"
)

type SalaryError struct {}

func (err *SalaryError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func throwSalaryError(salary int) (err error) {
	if salary < 150_000 {
		err = &SalaryError{}
		os.Exit(1)
	}else{
		fmt.Println("Debe pagar impuesto")
	}
	return 
}

func main()  {
	
	salary := 160_000

	err := throwSalaryError(salary)

	if err != nil {
		fmt.Println(err)
	}

	salary2 := 140_000

	//sin usar la funcion
	err1 := SalaryError{}

	if salary2 < 150_000{
		fmt.Println(err1.Error())
		os.Exit(1)
	}else{
		fmt.Println("Debe pagar impuesto")
	}


}