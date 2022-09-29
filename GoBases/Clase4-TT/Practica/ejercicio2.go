/*El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
 	Legajo
	Nombre y Apellido
	DNI
	Número de teléfono
	Domicilio

- Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos. Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”. Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.
- Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe. Para ello, necesitas leer los datos de un archivo .txt. En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt” (como en el ejercicio anterior, este archivo no existe, por lo que la función que intente leerlo devolverá un error). Debes manipular adecuadamente ese error como hemos visto hasta aquí.
Ese error deberá:
1.-   generar un panic;
2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmente.
- Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero (recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).
- Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos abiertos” (en ese orden). Utiliza defer para cumplir con este requerimiento.

Requerimientos generales:
- Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
- Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error (por ejemplo las que intenten leer archivos).
Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello (realiza también la validación pertinente para el caso de error retornado). */

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

/* --------------------------------- Tarea 1 -------------------------------- */

var NoIDError = errors.New("error: the file number can't be 0")
var EmptyFieldError = errors.New("error: the field can't be null")

type client struct{
	//legajo
	FileNumber int
	Name string
	DNI int
	PhoneNumber int
	Address string
}

func (cl *client) generateID() (err error) {
	randomID := rand.Int()
	if randomID == 0 {
		err = NoIDError
	} else {
		cl.FileNumber = randomID
	}

	return
}

/* --------------------------------- Tarea 2 -------------------------------- */
func readCustomersFile() (data string, err error) {

	receivedData, err := os.ReadFile("./customers.txt")
	if err != nil {
		err = errors.New("the indicated file was not found or is damaged")
	}

	data = string(receivedData)

	return
}


func checkClientExistence(clientDNI int, data string) bool {
	clients := strings.Split(data, "|")
	flag := false

	 for _, line := range clients {
		clientRefactor := strings.Split(line, ";")
		if clientRefactor[2] == fmt.Sprint(clientDNI) {
			flag = true
		}
		
	} 

	return flag


}

//llama a la funcion de existencia - el problema con esta funcion es que write file no agrega al archivo sino que lo pisa
/* func saveClient(cl client) (err error) {

	if !checkClientExistence(cl.DNI) {
		//registramos el cliente
		newClient := fmt.Sprintf("|\n%d;%s;%d;%d;%s", cl.FileNumber, cl.Name, cl.DNI, cl.PhoneNumber, cl.Address)
		receivedErr := os.WriteFile("./customers.txt", []byte(newClient), 0644)
		if receivedErr != nil{
			err = receivedErr
		}
		fmt.Println("The client was saved successfully")
	}else{
		fmt.Println("the client already exists")
	}

	return
} */

/* --------------------------------- Tarea 3 -------------------------------- */
func checkClientData(cl client) (correct bool, err error)  {
	correct = true
	if cl.FileNumber == 0 || 
		cl.Name == "" ||
		cl.DNI == 0 ||
		cl.PhoneNumber == 0 ||
		cl.Address == "" {
			err = EmptyFieldError
			correct = false
		}
	return
}

/* --------------------------------- Tarea 4 -------------------------------- */

func main()  {

	//siempre es mejor que las funciones retornen error y que en el main se maneje el panic

	countPanics := 0

	defer func ()  {
		if err := recover(); err != nil {
			fmt.Printf("An error ocurred: %v \n", err)
		}
		fmt.Println("\nExecution finished")

		if(countPanics > 0) {
			fmt.Printf("%d errors ocurred during execution\n", countPanics)
		}

		fmt.Println("There are no opened files")
		
	}()

	client := client{
		Name: "Camila Conte 2",
		DNI: 1111112,
		PhoneNumber: 12345,
		Address: "Street 1 567",
	}

	//tarea 1 - manejo de error

	err := client.generateID()
	if err != nil {
		panic(err)
	}


	//tarea 2 - leemos el archivo
	data, err1 := readCustomersFile()
	if err1 != nil {
		countPanics += 1
		panic(err)
	}else{
		//tarea 2 - registramos el cliente si no existe previamente
		clientExists := checkClientExistence(client.DNI, data)
		fmt.Println("Client exists : ", clientExists)

		//tarea 3 - validamos que los campos no sean nulos
		isCorrect, err := checkClientData(client)
		fmt.Println("Client is correct : ", isCorrect)
		if err != nil {
			countPanics += 1
			panic(err)
		}


 	}


}