package main

import "fmt"

func main()  {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//para declarar un array con valores predeterminados
	var myArray = [...]int{1,2,3,4,5,6}
	fmt.Println(myArray)

	//recorte entre la posicion 2 a la 4 (no incluida)
	slc := myArray[2:4]

	//recorte entre 0 y 3 (no incluido)
	slc1 := myArray[:3]

	//recorte de principio a fin
	slc2 := myArray[:]
	slc3 := myArray[0:6]

	fmt.Println(slc)
	fmt.Println(slc1)
	fmt.Println(slc2)
	fmt.Println(slc3)

	//vemos la longitud y capacidad de un slice (tambien se puede hacer con un array)
	fmt.Println(len(slc3), cap(slc3))
	

	//agregamos elementos a un slice
	slc3 = append(slc3, 7,8,9,1)

	fmt.Println(slc3)
	fmt.Println(len(slc3), cap(slc3))

	//la definicion que utilizaremos será 
	//var s []int
}