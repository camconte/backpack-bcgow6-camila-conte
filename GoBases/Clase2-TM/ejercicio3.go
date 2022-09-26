/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/
package main

import "fmt"

func calcularSalario(minutosPorMes int, categoria string) (salario float64) {
	horasPorMes := minutosPorMes/60

	var porcentajeSalario float64
	switch categoria {
	case "C":
		salario = float64(horasPorMes*1000)
	case "B":
		salario = float64(horasPorMes*1500)
		porcentajeSalario = salario*0.20
		salario += porcentajeSalario
	case "A":
		salario = float64(horasPorMes*3000)
		porcentajeSalario = salario*0.50
		salario += porcentajeSalario
	}
	return
}

func main(){
	fmt.Println(calcularSalario(3600, "A"))	
}