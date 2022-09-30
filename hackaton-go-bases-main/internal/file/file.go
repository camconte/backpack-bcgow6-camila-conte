package file

//modulo encargado de leer el archivo y manipularlo
import (
	//"fmt"
	"fmt"
	"hackathon/internal/service"
	"os"

	"github.com/gocarina/gocsv"
)

//encoding/csv
//newReader pide los bytes de la data


type File struct {
	Path string `csv:"Path"`
}

func (f *File) Read() (tickets []service.Ticket, err error) {

	dataBytes, err0 := os.ReadFile(f.Path)
	if err0 != nil{
		err = err0
	}


	err1 := gocsv.UnmarshalBytes(dataBytes, &tickets)
	if err1 != nil {
		err = err1
	}

	return
}

func (f *File) Write(ticket service.Ticket) (err error) {

	ticketsFile, err := os.OpenFile(f.Path, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer ticketsFile.Close()


	//leemos el archivo
	dataBytes, err0 := os.ReadFile(f.Path)
	if err0 != nil{
		err = err0
	}

	var tickets []service.Ticket

	err1 := gocsv.UnmarshalBytes(dataBytes, &tickets)
	if err1 != nil {
		err = err1
	}

	stringTicket := fmt.Sprintf(
		"\n%d,%s,%s,%s,%s,%d",
		ticket.Id,
		ticket.Names,
		ticket.Email,
		ticket.Destination,
		ticket.Date,
		ticket.Price,
	)


    _, err2 := ticketsFile.Write([]byte( stringTicket ))
    if err2 != nil {
        err = err2
    }

	for i, ticketInFile := range tickets {
		if ticketInFile.Id == ticket.Id {
			tickets[i] = ticket
		}
	}

	ticketsString, err2 := gocsv.MarshalString(tickets) 
	if err2 != nil {
		err = err2
	}

	err3 := os.WriteFile(f.Path, []byte(ticketsString), 0644)
	if err3 != nil {
		err = err3
	}

    ticketsFile.Close()

	return 
}
