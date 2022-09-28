package main

import (
	"errors"
	"fmt"
	//"os"
)

//creamos una estructura
type MyCustomError struct{
	StatusCode int
	Message string
}
//implementamos el Error para que sea un error y cumpla con el "contrato"
func (err *MyCustomError) Error() string  {
	return fmt.Sprintf("%s (%d)", err.Message, err.StatusCode)
}

/* func obtainError(status int) (code int, err error) {
	if status >= 400 {
		return 500, &MyCustomError{
			StatusCode: 500,
			Message: "Algo salió mal",
		}
	}
	return 200, nil
} */


var errTest = errors.New("error Test 1")

func getError() error {
	return fmt.Errorf("informacion extra: %w", errTest)
}

type ErrType2 struct{}

func (e ErrType2) Error() string {
	return "soy el error 2"
}

func main()  {
	//status, err := obtainError(400)

	/* status := 200

	if status > 399{
		//por convencion los errores empiezan con minuscula y no deben tener un punto al final
		err := errors.New("la petición ha fallado")
		fmt.Println(err)
		os.Exit(1)
	} */

	/* if err != nil{
		fmt.Println(err)
		//si el codigo es distinto de 0 significa que el programa cerró el con un error
		//exit codes: https://tldp.org/LDP/abs/html/exitcodes.html
		os.Exit(1)
	} */

	//fmt.Printf("Status %d, funciona!\n", status)

	//ambos comparten el mismo tipo y estructura
	err1 := &MyCustomError{
		StatusCode: 502,
		Message: "Soy el error 1",
	}

	err2 := &MyCustomError{
		StatusCode: 406,
		Message: "Soy el error 2",
	}

	//son iguales?
	//si no le ponemos el & estamos comparando de manera directa dos errores, lo cual no se puede
	//errors.As se encarga de chequar si un error es de un tipo especifico, en este caso si es un MyCustomError
	bothErrorsAreEqual := errors.As(err1, &err2)

	fmt.Println(bothErrorsAreEqual)

	//errors.Is compara un error con un valor, un error coincide con un objetivo si es igual a este

	err := getError()

	//compara el err con el errTest (valor)
	//conviene usar este metodo en lugar de == : https://go.dev/blog/go1.13-errors
	coincidence := errors.Is(err, errTest)

	fmt.Println(coincidence)

	//Unwrap para desencadenar errores
	err4 := ErrType2{}
	err3 := fmt.Errorf("soy el error 3, contengo al 4 %w", err4)

	fmt.Println(
		errors.Unwrap(err3),
	)

	fmt.Println(
		errors.Unwrap(err4),
	)

	//SIEMPRE VALIDAR LOS ERRORES, NUNCA IGNORARLOS
}