package file

//modulo encargado de leer el archivo y manipularlo
import (
	//"fmt"
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
	//leer el archivo donde se encuentran los tickets del dia
	/* data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, errors.New("the indicated file was not found or is damaged")
	}

	stringData := string(data)

	stringTickets := strings.Split(stringData, "\n")

	 for _, line := range stringTickets {
		ticketRefactor := strings.Split(line, ",")
		
		ticketID, err0 := strconv.ParseInt(ticketRefactor[0], 0, 24)
		ticketNames := ticketRefactor[1]
		ticketEmail := ticketRefactor[2]
		ticketDestination := ticketRefactor[3] 
		ticketDate := ticketRefactor[4]
		ticketPrice, err1 := strconv.ParseInt(ticketRefactor[5], 0, 24)
		if err1 != nil{
			err = err1
			return 
		}else if err0 != nil {
			err = err0
		}
		

		ticket := service.Ticket{
			Id: int(ticketID),
			Names: ticketNames,
			Email: ticketEmail,
			Destination: ticketDestination,
			Date: ticketDate,
			Price: int(ticketPrice),
		}

		tickets = append(tickets, ticket)
	} 
 */

	

	dataBytes, err0 := os.ReadFile(f.Path)
	if err0 != nil{
		err = err0
	}


	err1 := gocsv.UnmarshalBytes(dataBytes, &tickets)
	if err1 != nil {
		err = err1
	}

	//tickets = dataTickets

	return
}

func (f *File) Write(ticket service.Ticket) (err error) {

/* 	ticketsFile, err := os.OpenFile("tickets2.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer ticketsFile.Close()

	//ver como actualizar el archivo
	//traemos lo que tiene el archivo
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return errors.New("the indicated file was not found or is damaged")
	}

	stringData := string(data)

	stringTickets := strings.Split(stringData, "\n")

	ticketId := fmt.Sprint(ticket.Id)
	ticketPrice := fmt.Sprint(ticket.Price)

	exists := false

	for _, ticketInFile := range stringTickets {
		ticketRefactor := strings.Split(ticketInFile, ",")
		//si ya existe el ticket, lo modificamos:
		if ticketRefactor[0] == fmt.Sprint(ticket.Id) {
			//primero lo vaciamos
			ticketRefactor = []string{}
			//luego le agregamos los datos nuevos
			ticketRefactor = append(ticketRefactor, ticketId, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticketPrice)
			//indicamos que si existe
			exists = true
		}
	}

	if !exists {
		//si, luego de recorrer el array, no encuentra ningun ticket con el mismo id, entonces lo agregamos
		newStringTicket := fmt.Sprint(ticketId, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticketPrice)
		stringTickets = append(stringTickets, newStringTicket)
		err := gocsv.MarshalFile(stringTickets, ticketsFile)
		if err != nil{
			panic(err)
		}
	} */

	ticketsFile, err := os.OpenFile("tickets2.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
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


	//chequeamos si el ticket ya existe o si es uno nuevo
	exists := false

	for i, ticketInFile := range tickets {
		if ticketInFile.Id == ticket.Id {
			exists = true

			//actualizamos los datos que correspondan
			switch {
			case ticketInFile.Names != ticket.Names:
				ticketInFile.Names = ticket.Names
			case ticketInFile.Email != ticket.Email:
				ticketInFile.Email = ticket.Email
			case ticketInFile.Destination != ticket.Destination:
				ticketInFile.Destination = ticket.Destination
			case ticketInFile.Date != ticket.Date:
				ticketInFile.Date = ticket.Date
			case ticketInFile.Price != ticket.Price:
				ticketInFile.Price = ticket.Price
			}

			tickets[i] = ticketInFile
		}
	}

	//si no existe lo agregamos
	if !exists {
		tickets = append(tickets, ticket)
	}

	//actualizamos el archivo
	err2 := gocsv.MarshalFile(&tickets, ticketsFile)
	if err2 != nil {
		err = err2
	}

	return 
}
