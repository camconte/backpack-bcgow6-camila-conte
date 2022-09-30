package main

import (
	"fmt"
	"hackathon/internal/file"
	"hackathon/internal/service"
)
func main() {
	defer func () {
		if err := recover(); err != nil {
		fmt.Printf("An error ocurred: %v \n", err)
		}
		fmt.Println("\nExecution finished")
	}()

	ticketsFile := file.File{
		Path: "./tickets.csv",
	} 
	
	var tickets []service.Ticket
	var err error
	
	
	//ejecutamos la funcion de Read y almacenamos en la variable tickets
	tickets, err = ticketsFile.Read()
	
	if err != nil {
		panic(err)
	} 
	
	bookings := service.NewBookings(tickets)

	/* --------------------------- creacion de ticket --------------------------- */

	//luego ejecutar write cada vez que sea necesario
	/* newTicket := service.Ticket{
		Names: "nuevoTicket ",
		Email: "pruebaFinal",
		Destination: "TEST",
		Date: "23:57",
		Price: 1500,
	} 

	//creamos el ticket en el booking y luego lo guardamos en el archivo el nuevo ticket
	ticket, err0 := bookings.Create(newTicket)
	if err0 != nil {
		panic(err0)
	}else{
		err01 := ticketsFile.Write(ticket)
		if err01 != nil {
			panic(err01)
		}
	} */

	/* --------------------------- busqueda de ticket --------------------------- */
	//buscamos un ticket
	/* foundTicket, err1 := bookings.Read(1001)
	if err1 != nil {
		panic(err1)
	}else{
		fmt.Println(foundTicket)
	} */

	/* ------------------------- actualizacion de ticket ------------------------ */

	updateTicket := service.Ticket{
		Id: 1001,
		Names: "PROBANDO TEST",
		Email: "camilatest.dev@gmail.com",
		Destination: "Spain",
		Date: "22:57",
		Price: 1800,
	}

	_, err2 := bookings.Update(updateTicket.Id, updateTicket)
	if err2 != nil{
		panic(err2)
	}else{
		err20 := ticketsFile.Write(updateTicket)
		if err20 != nil {
			panic(err20)
		}
	}

	/* ------------------------ eliminacion de un ticket en memoria ------------------------ */
	/* //probamos el delete
	toDeleteId := 1

	//buscamos el ticket a eliminar
	toDeleteTicket, err3 := bookings.Read(toDeleteId)
	if err3 != nil {
		panic(err3)
	}else{
		_, err4 := bookings.Delete(toDeleteTicket.Id)
		if err4 != nil{
			panic(err4)
		}
	} */


	
}
