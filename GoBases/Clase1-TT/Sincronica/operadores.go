package main

import "fmt"

func main()  {
	//operadores aritmeticos
	x, y, q := 10, 20, 30

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)

	//operadores de asignación
	var z, p = 15, 25
	z = p
	fmt.Println("= ", z)
	z = 15
	z += p
	fmt.Println("+=", z)
	z = 50
	z -= p
	fmt.Println("-=", z)
	z = 2
	z *= p
	fmt.Println("*=", z)
	z = 100
	z /= p
	fmt.Println("/=", z)
	z = 40
	z %= p
	fmt.Println("%=", z)

	//operadores de comparación
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
	fmt.Println(x < y)
	fmt.Println(x <= y)

	//operadores lógicos
	fmt.Println(x < y && x > q)
	fmt.Println(x < y || x > q)
	fmt.Println(!(x == y && x > q))

	//operadores de dirección
	text := "Hola Mundo"
	var pText *string

	pText = &text

	fmt.Println(pText) /* 0xc0000a0a0 */
	fmt.Println(*pText) /* Hola Mundo */

}
