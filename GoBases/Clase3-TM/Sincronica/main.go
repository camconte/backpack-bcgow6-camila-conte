package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	messageHello := "Hello World :)"
	messageBye := "Bye World :("

	fmt.Print(messageHello, messageBye)

	fmt.Println(messageHello)
	fmt.Println(messageBye)

	flag := true
	flag2 := false
	fmt.Printf("La variable es %t y la otra es %t\n", flag, flag2) 

	number := 12222.221314

	fmt.Printf("El numero resumido es %.3f\n", number)

	nombre := "Camila"
	edad := 20
	
	fmt.Printf("Hola %s! Tienes %d años\n", nombre, edad)

	fmt.Printf("Hola %10s! Tienes %d años\n", nombre, edad)

	//sprint nos va a retornar la cadena generada
	stringGenerada := fmt.Sprint("Hola ", nombre, "! Tienes ", edad, " años")
	fmt.Println(stringGenerada)
	
	stringGenerada2 := fmt.Sprintf("Hola %s! Tienes %d años", nombre, edad)
	fmt.Println(stringGenerada2)

	//package os

	//creamos la variable
	err := os.Setenv("NAME", "Camila")
	if err != nil{
		panic(err)
	}

	//obtenemos la variable
	name := os.Getenv("NAME")
	fmt.Println(name)

	//llega un string vacio porque no existe la misma
	name2 := os.Getenv("jasnenfjdnf")
	fmt.Println(name2)

	//para chequear que existe la variable de entorno
	value, ok := os.LookupEnv("USERPROFILE")
	fmt.Println(value, ok)

	//exit tiene distintos status code
	bandera := true

	if bandera {
		fmt.Println("flag is true")
	}else{
		fmt.Println("flag es false")
		os.Exit(1)
	}

	fmt.Println("Fin de instrucciones")

	//fileContentAsBytes, err := os.ReadFile("./main.go")

	//fmt.Printf("%s", fileContentAsBytes)

	//message := "Hello World, I'm in a file"

	//Crea el archivo si no existe
	/*err1 := os.WriteFile("./test.txt", []byte(message), 0644)

	if err1 != nil{
		panic(err1)
	}else{
		fmt.Println("File Written successfully!")
	}*/

	//package io

	//copiar desde un reader a una salida estandar
	//un reader es una interfaz, un contrato de metodos que define como se va a interacturar con esa estructura de datos 
	reader := strings.NewReader("Hello Reader!")

	//la salida estandar los va a representar en la terminal
	_, err3 := io.Copy(os.Stdout, reader)

	if err3 != nil {
		panic(err3)
	}

	//Vamos a leer el contenido del reader
	b, err4 := io.ReadAll(reader)
	if err4 != nil {
		panic(err4)
	}

	fmt.Println(string(b))

	//la salida estandar (os.Stdout) se puede reemplazar por un log por ejemplo
	io.WriteString(os.Stdout, "Hello World!")
	
}