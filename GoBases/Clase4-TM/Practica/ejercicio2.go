/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

package main

import (
	"fmt"
	"errors"
	"os"
)

var SalaryError = errors.New("error: el salario ingresado no alcanza el mínimo imponible")

func throwSalaryError(salary int) (err error) {
	if salary < 150_000 {
		err = SalaryError
		os.Exit(1)
	}else{
		fmt.Println("Debe pagar impuesto")
	}
	return 
}

func main()  {
	
	salary := 170_000

	err := throwSalaryError(salary)

	if err != nil {
		fmt.Println(err)
	}

	salary2 := 140_000

	//sin usar la funcion
	err1 := SalaryError

	if salary2 < 150_000{
		fmt.Println(err1.Error())
		os.Exit(1)
	}else{
		fmt.Println("Debe pagar impuesto")
	}

}