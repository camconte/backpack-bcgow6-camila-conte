/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/
package main

import "fmt"

type Student struct {
	Name string
	Lastname string
	DNI string
	Date string
}

func (s Student) detail(){
	fmt.Printf("Student: %+v \n", s)
}

func main()  {
	student := Student{
		Name: "Camila",
		Lastname: "Conte",
		DNI: "00.000.000",
		Date: "11-06-2002",
	}

	student.detail()
}