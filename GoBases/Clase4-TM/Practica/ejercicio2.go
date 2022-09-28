/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

package main

import (
	"fmt"
	"errors"
)

var SalaryError = errors.New("error: el salario ingresado no alcanza el mínimo imponible")

func throwSalaryError(salary int) (err error) {
	if salary < 150_000 {
		err = SalaryError
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

}