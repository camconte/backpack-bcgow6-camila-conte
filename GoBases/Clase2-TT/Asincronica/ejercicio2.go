/*Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:

- Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
- Print: Imprime por pantalla la matriz de una formas más visible (Con los   	saltos de línea entre filas)

La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main

import "fmt"

type Matrix struct {
	Values []float64
	Height int
	Width int
	Quadratic bool
	MaxValue float64
}

func (m *Matrix) setValues(values ...float64) {
	m.Values = values
}

func (m Matrix) printMatrix(){
	for i := 0; i < len(m.Values); i++ {
		if (i+1)%m.Width == 0{
			fmt.Println(m.Values[i])
		}else{
			fmt.Printf("%v	", m.Values[i])
		}
		
	}
}

func main(){
	m := Matrix{
		Height: 3,
		Width: 4,
		Quadratic: true,
	}

	m.setValues(2, 1, 45, 23, 1, 8, 10, 5, 2, 5, 2, 3, 4)
	m.printMatrix()
}