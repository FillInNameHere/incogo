package domain

type Address struct {
	ID      uint `gorm:"primary_key"`
	BPRefer uint
	UUID    string
	Street  string
	Number  string
	Zip     string
	City    string
	State   string
	Country string
}

func NewAddress(street string) *Address {
	return &Address{Street: street}
}

func (a *Address) SetStreet(street string) {
	a.Street = street
}

func (a *Address) SetNumber(number string) {
	a.Number = number
}

func (a *Address) SetZip(zip string) {
	a.Zip = zip
}

func (a *Address) SetCity(city string) {
	a.City = city
}

func (a *Address) SetState(state string) {
	a.State = state
}

func (a *Address) SetCountry(country string) {
	a.Country = country
}