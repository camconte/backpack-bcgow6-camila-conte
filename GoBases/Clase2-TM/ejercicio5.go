/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

ejemplo:
const (

	dog = "dog"
	cat = "cat"

)

...

animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)

...

var amount float64
amount+= animaldog(5)
amount+= animalCat(8)
*/
package main

import (
	"errors"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	tarantula = "tarantula"
 )

func dogFoodQuantity(quantity int) (totalFood float64){
	totalFood = float64(quantity)*10
	return
}
func catFoodQuantity(quantity int) (totalFood float64){
	totalFood = float64(quantity)*5
	return
}
func hamsterFoodQuantity(quantity int) (totalFood float64){
	totalFood = float64(quantity)*0.250
	return
}
func tarantulaFoodQuantity(quantity int) (totalFood float64){
	totalFood = float64(quantity)*0.150
	return
}


func animal(animal string) (function func(int) float64, err error){
	switch animal {
	case dog:
		function = dogFoodQuantity
	case cat:
		function = catFoodQuantity
	case hamster:
		function = hamsterFoodQuantity
	case tarantula:
		function = tarantulaFoodQuantity
	default:
		err = errors.New("El animal no existe")
	}

	return
}

func main()  {
	animalDog, err := animal(dog)
	animalCat, err := animal(cat)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)

	if err != nil {
		fmt.Println(err.Error())
	}else{
		var amount float64
		amount += animalDog(2)
		amount += animalCat(1)
		amount += animalHamster(2)
		amount += animalTarantula(3)
		fmt.Println(amount)
	}

	

	fmt.Println()
}