package service

import (
	"errors"
	"strconv"
	"strings"
)

var ErrTime = errors.New("error: the ticket time limit is 23:59")
var ErrNotFoundTicket = errors.New("error: no ticket was found with the id received")
var ErrAlreadyExist = errors.New("error: the id received is already in use")

//consume el file para actualizar los valores
type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id int `csv:"Id"`
	Names string `csv:"Names"`
	Email string `csv:"Email"`
	Destination string `csv:"Destination"`
	Date string `csv:"Date"`
	Price int `csv:"Price"`
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (newTicket Ticket, err error) {
	

	for _, ticket := range b.Tickets {
		if ticket.Id == t.Id{
			return Ticket{}, ErrAlreadyExist
		}
	}

	ticketTime := strings.Split(t.Date, ":")
	hourTicket, err1 := strconv.ParseInt(ticketTime[0], 0, 64)
	minuteTicket, err0 := strconv.ParseInt(ticketTime[1], 0, 64)

	if err0 != nil{
		err = err0
	}
	if err1 != nil{
		err = err1
	}
		
	if hourTicket > 23 || minuteTicket > 59{
		err = ErrTime
	} else {
		t.Id = b.Tickets[len(b.Tickets)-1].Id + 1
		b.Tickets = append(b.Tickets, t)

		b.Tickets = append(b.Tickets, t)
		
		newTicket = t
	}


	return
}

func (b *bookings) Read(id int) (t Ticket, err error) {
	flag := false

	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			t = ticket

			//indicamos que el ticket fue encontrado
			flag = true
		}
	}

	if !flag {
		err = ErrNotFoundTicket
	}

	return 
}

func (b *bookings) Update(id int, t Ticket) (ticket Ticket, err error) {
	exists := false

	for i, ticketInBookings := range b.Tickets {
		if ticketInBookings.Id == id {
			
			exists = true

			b.Tickets[i] = t 
			ticket = t
			
		}
	}


	if !exists {
		err = ErrNotFoundTicket
	}

	return
}

func (b *bookings) Delete(id int) (deleteId int, err error) {
	for i, ticket := range b.Tickets {
		if ticket.Id == id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i + 1:]...)
			deleteId = id
		}
	}

	if deleteId == 0 {
		err = ErrNotFoundTicket
	}

	return 
}
