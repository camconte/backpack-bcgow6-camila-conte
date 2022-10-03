package main

import (
	"fmt"
	"time"
)
//Recibimos el channel tambien
func procesar(i int, ch chan int)  {
	fmt.Println(i, "- Inicia")
	time.Sleep(1000*time.Millisecond)
	fmt.Println(i, "- Termina")

	//escribimos sobre el canal
	ch <- i
}

//os canales son una manera segura de trabajar con concurrencia, a diferencia de lo que es memoria compartida. Los canales buscan pasar datos entre distintos pipes. Unos dejan y otros toman. Y nosotros diseñamos el sistema

func main()  {
	//solo se define un solo canal y todos van a consumir del mismo
	//Definimos el channel el cual leeremos para coordinar las go routines
	canal := make(chan int)

	now := time.Now()

	fmt.Println("Inicia el programa")

	for i := 0; i < 10; i++ {
		go procesar(i, canal)
		
	}

	//necesitamos leer todas las entradas al canal
	for i := 0; i < 10; i++ {
		//hasta que no se lee el canal, no continua la ejecucion
		//leemos de canal y lo asignamos a una variable
		lectura := <- canal
		//ahora leemos la variable
		fmt.Println("Lectura de canal: ", lectura)
		
	}

	fmt.Println("Termina el programa")
	fmt.Printf("El programa demoró: %d ms\n", time.Now().Sub(now).Milliseconds())
}