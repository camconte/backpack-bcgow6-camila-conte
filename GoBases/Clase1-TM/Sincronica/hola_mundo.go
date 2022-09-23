package main

import "fmt"

// CONSTANTES - se puede obviar el tipo de dato
const SEGUNDOS = 0

// var horas int
func main() {
	//horas = 25
	//horas := 24
	//fmt.Println(horas)
	//fmt.Printf("%T", float64(horas)) //sirve para formatear la salida y, de esta manera, imprimimos el tipo de la variable
	//no se mantiene la conversion del tipo de datos en la variable ya que no lo volvimos a almacenar
	//fmt.Printf("%T", horas)
	fmt.Printf("%T", SEGUNDOS)
}

//comando para inicializar el modulo
//go mod init github.com/camconte/backpack-bcgow6-camila-conte

//comando para agregar una dependencia, en este caso, la dependencia de gin gonic
//go get -u github.com/gin-gonic/gin
