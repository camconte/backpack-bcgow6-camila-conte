/*Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
*/
package main

import "fmt"

func calcularImpuestosSalario(salario float64) float64{
    salarioCalculado := salario

    if salario > 50000 && salario <= 150000{
        salarioCalculado -= salario*0.17
    }else if salario > 150000 {
        salarioCalculado -= salario*0.27
    }


    return salarioCalculado
}

func main() {

    fmt.Println(calcularImpuestosSalario(40000))
}