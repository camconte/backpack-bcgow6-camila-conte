package main

import (
	"fmt"
	"hackathon/internal/file"
	"hackathon/internal/service"
)
func main() {
	
	updateTicket := service.Ticket{
		Id: 5,
		Names: "HOLAAA Nobes",
		Email: "snobes4@google.com.au",
		Destination: "Czech Republic",
		Date: "0:31",
		Price: 1398}

	ticketsFile := file.File{
		Path: "./tickets2.csv",
	} 

	//ticketsFile.Write(newTicket)
	ticketsArray, err := ticketsFile.Read()
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println(ticketsArray)

	ticketsFile.Write(updateTicket)
	
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
