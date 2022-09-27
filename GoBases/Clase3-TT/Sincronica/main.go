package main

import "fmt"

func main()  {
	//puntero -> accedemos a la direccion que este almacenada en él
	//&puntero -> accedemos a la direccion de memoria del puntero
	//*puntero -> accedemos al valor que haya almacenado en la direccion de memoria almacenada en el puntero. Por ejemplo: si tenemos almacenada la direccion de memoria de una variable x que vale 20, entonces accederemos al valor 20

	//&variable -> vemos la direccion de memoria de la variable
	//*variable -> referenciamos a la variable "original" y no una copia
	//definicion de punteros
	//manera 1
	//var puntero *int

	//manera2
	//puntero2 := new(int)
	
	//manera3
	var numero int
	//obtenemos la direccion de memoria de la variable numero
	p3 := &numero

	//devuelven la misma direccion de memoria que es: la direccion de memoria de la variable numero
	fmt.Printf("La direccion es: %p\n", p3)
	fmt.Printf("La direccion es: %p\n", &numero)

	//desreferenciacion
	//supongamos que queremos modificar la variable a traves del puntero
	fmt.Printf("El valor del numero es %d\n", numero)
	
	//si no ponemos el * no vamos a modificar la variable ya que estamos modificando una "copia"
	*p3 = 20
	//una de las utilidades:
	Incrementar(p3)

	prueba1 := Prueba{Valor:10}

	prueba1.Actualizar(12)

	//el caracter de escape +v nos imprime más detalle de la variable, en este caso, de la estructura
	fmt.Printf("El valor de la prueba es %+v\n", prueba1)
	
	
	fmt.Printf("El valor del numero es %d\n", numero)
}

func Incrementar(puntero *int){
	*puntero = 30
}

type Prueba struct{
	Valor int
}

//si no ponemos el * no actualizaremos el valor original sino la copia
func (p *Prueba) Actualizar (new int){
	p.Valor = new
}