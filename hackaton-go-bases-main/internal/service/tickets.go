package service

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

func (b *bookings) Create(t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
