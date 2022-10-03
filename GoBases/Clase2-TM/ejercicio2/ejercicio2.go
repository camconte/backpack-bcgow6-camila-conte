/*Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/

package main

import (
	"fmt"
	"errors"
)

func calcularPromedio(calificaciones ...int) (promedio int, err error){

	for _, calificacion := range calificaciones {
		if calificacion >= 0 {
			promedio += calificacion
		}else{
			return 0, errors.New("Las calificaciones deben ser números mayores a cero")
		}
	}

	promedio /= len(calificaciones)

	return 
}

func main(){
	promedio, err := calcularPromedio(10, 4, 6, 7, 9, 7)

	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(promedio)
	}
}