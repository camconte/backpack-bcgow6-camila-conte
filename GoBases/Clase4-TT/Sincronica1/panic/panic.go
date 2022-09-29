package main

import (
	"fmt"
	"os"
)

func main()  {
	//Se ejecuta al final de la ejecucion de la funcion que la invoca
	//Es decir, se ejecuta al finalizar el main
	//puede ser una funcion anonima o no
	defer func ()  {
		//Recover se debe usar dentro de una funcion diferida si o si
		//esta funcion toma el mensaje de panic y lo transforma en un error, por lo tanto, lo podemos tratar como tal
		if err := recover(); err != nil{
			fmt.Println("ocurrió un panic: ", err)
		}
		fmt.Println("Hola estoy diferida")
	}()

	//situacion de panic
	_, err := os.ReadFile("este archivo no existe")
	if err != nil {
		panic(err)
	}



	fmt.Println("Todo correcto")




}