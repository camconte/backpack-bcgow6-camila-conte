package main

import (
	//"encoding/json"
	"fmt"
	//"reflect"
)

//cuando se inicializan estan vacios sus campos
//en cuanto a nomenclatura, lo mismo sucede con el nombre de la estructura:
//si esta en minuscula va a ser privada para el paquete
//por lo general, se usan privadas las estructuras
type Perro struct {
	//si el nombre del campo se escribe en minuscula, significa que no va a poder ser accedido. El campo va a ser privado para el paquete main, en este caso
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Color [2]string `json:"color"`
	//campo privado - el paquete encoding json no lo puede leer
	marca string //`json:"marca"`
	//tiene un campo que apunta a un nuevo tipo de dato: Duenio
	//estructura compuesta por otra estructura
	Owner Duenio `json:"owner"`
}

func (p Perro) Ladrar(){
	fmt.Println("Guau")
}

func (p Perro) Saludar(){
	fmt.Println("Guau! Soy ", p.Nombre)
}
//no se cambia el nombre en el objeto dog, sino en una copia
/*func (p Perro) Renombrar(newName string){
	p.Nombre = newName
	fmt.Println("Guau! Ahora soy ", p.Nombre)
}*/

//con ese agregado en la firma, se modifica el objeto en si cuando llamemos al metodo
//esto es el agregado de punteros
func (p *Perro) Renombrar(newName string){
	p.Nombre = newName
	fmt.Println("Guau! Ahora soy ", p.Nombre)
}

func (p Perro) Dormir(){
	fmt.Println("Estoy durmiendo")
}

func (p Perro) Comer(){
	fmt.Println("Estoy comiendo")
}

type Duenio struct {
	Documento int
	Nombre string
	NroDeContacto int
}

type Gato struct {
	Nombre string
	Edad int
	Raza string
	Peso float64
	Caracter string
}

func (g Gato) Dormir(){
	fmt.Println("Estoy durmiendo")
}

func (g Gato) Comer(){
	fmt.Println("Estoy comiendo")
}

type Animal interface {
	Dormir()
	Comer()
}

//inicializador de animal de tipo perro
func NewAnimal(tipo string) Animal {
	//retorna una referencia al tipo que va a implementar la interfaz.
	//este tipo de dato previamente debe contar con los metodos que se implementan en la interfaz
	if tipo == "Perro"{
		return &Perro{}
	}else if tipo == "Gato"{
		return &Gato{}
	}
	return nil
}




func main()  {
	//para inicializar los valores:
	/*dog := Perro{
		//manera 1
		"Firulais", 2, "Pastor Alemán", 10, [2]string{"Negro", "Blanco"}, "prueba",
	} */


	//si no sabemos el valor de algun campo de la estructura, utilizamos la otra manera:
	dog1 := Perro{
		Nombre: "Firulais Jr.", 
		Edad: 3, 
		Raza: "Pastor Alemán", 
		Peso: 12,
		Owner: Duenio{
			Documento: 12345,
			Nombre: "Carlos",
			NroDeContacto: 1234567,
		},
	}

	//Metodos
	dog1.Ladrar()
	dog1.Saludar()
	
	dog1.Renombrar("Cacho")

	//creamos un perro que ya implementa la interfaz, desde la "fabrica" de animales
	nuevoPerro := NewAnimal("Perro")
	nuevoPerro.Comer()
	nuevoPerro.Dormir() 

	//fmt.Printf("Dog: %+v \n", dog)
	//fmt.Printf("Dog: %+v \n", dog1)

	//actualizamos un campo
	//dog1.Owner.Documento = 1234

	//fmt.Printf("Owner.Documento: %+v \n", dog1.Owner.Documento)

	//se hace un cifrado a json en donde se utilizan las etiquetas definidas en la estructura
	//el operador & nos da la direccion de memoria
	/*b, err := json.Marshal(&dog1)

	if err != nil {
		fmt.Println("Hubo un error")
	}else {
		fmt.Printf("\njson.Marshal: %+v\n", string(b))
	}*/

	//ejemplo de reflect : paquete que se suele utilizar para serializaciones customizadas. Con él se puede ver la etiqueta asociada a cada campo
	//https://pkg.go.dev/reflect

	/*dogR := reflect.TypeOf(dog1)

	for i := 0; i < dogR.NumField(); i++ {
		field := dogR.Field(i)
		fmt.Printf("\nCampo: %s, tiene %s\n", field.Name, field.Tag.Get("json"))
	}*/

}

