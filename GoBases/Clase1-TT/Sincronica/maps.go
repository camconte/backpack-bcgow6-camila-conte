package main

import "fmt"

func main()  {
	var mapa = map[string]int{"Cami":20, "Cris":26}

	//insertamos un nuevo elemento
	mapa["Hola"] = 10

	//extraemos un valor de un map y lo almacenamos en una variable
	//el ok va a ser un booleano y nos va a indicar si el elemento se encontro en el map o no
	x, ok := mapa["Hola"] //nos devuelve true

	//para ignorar el almacenamiento en una variable y solo preguntar si existe o no el elemento con la clave definida
	_, ok1 := mapa["Mila"]

	fmt.Println(mapa)
	fmt.Println(mapa["Cami"])
	fmt.Println(x, ok)
	fmt.Println(ok1)
}