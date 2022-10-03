/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

# Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

Ejemplo:

const (

	minimum = "minimum"
	average = "average"
	maximum = "maximum"

)

...

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...

minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/
package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )

func minimumFunction(nums ...int) (min int) {
	min = nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return
}

func averageFunction(nums ...int) (average int){
	
	for _, num := range nums {
		average += num
	}

	average /= len(nums)
	
	return average
}

func maximumFunction(nums ...int) (max int) {
	max = nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	
	return
}


func operation(operation string) (func (...int) int, error){
	switch operation {
	case minimum:
		return minimumFunction, nil
	case average:
		return averageFunction, nil
	case maximum:
		return maximumFunction, nil
	default:
		return nil, errors.New("The operation is not defined")
	}
}

func main()  {
	minFunction, err := operation(minimum)
	avFunction, err := operation(average)
	maxFunction, err := operation(maximum)

	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(minFunction(1, 2, 3, 4, 5))
		fmt.Println(avFunction(1, 2, 3, 4, 5))
		fmt.Println(maxFunction(1, 2, 3, 4, 5))
	}
}