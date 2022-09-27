package main

import (
	"fmt"
	"time"
)

func procesar(i int)  {
	fmt.Println(i, "- Inicia")
	//simulamos una demora
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "- Termina")
}

func main()  {
	now := time.Now()
	fmt.Println("Inicia el programa")
	//se ejecuta 1 instruccion atras de otra
	for i := 0; i < 10; i++ {
		
		//procesar(i)
		//se llamaron los procesos y siguio la ejecucion, es decir, no espero a que terminen de ejecutarse para finalizar el programa
		go procesar(i)
	}
	//le agregamos un sleep para darle tiempo a la funcion a imprimir por consola
	//el tiempo que indicamos de espera es el tiempo aproximado en el que le decimos al programa que debe ejecutarse -	esto no es lo recomendable, para eso se usan los channels
	time.Sleep(2000*time.Millisecond)
	fmt.Println("Termina el programa")
	fmt.Printf("El programa demoró: %d ms\n", time.Now().Sub(now).Milliseconds())
}