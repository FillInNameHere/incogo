package domain

import (
	//"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/database"
)

type Businesspartner struct {
	//database.Model
	ID 		   uint `gorm:"primary_key"`
	Name           string
	Private        bool
	Contactpersons []Contactperson `gorm:"ForeignKey:BPRefer"`
	Addresses      []Address `gorm:"ForeignKey:BPRefer"`
}

func NewBusinesspartner(name string, private bool) *Businesspartner {
	return &Businesspartner{Name: name, Private: private}
}

func (b *Businesspartner) SetName(name string) {
	b.Name = name
}

func (b *Businesspartner) SetPrivate(private bool) {
	b.Private = private
}

/*func (b *Businesspartner) SetAddresses(addresses []Address) {
	b.Addresses = addresses
}

func (b *Businesspartner) AddAddress(address Address) {
	b.Addresses = append(b.Addresses, address)
}

func (b *Businesspartner) SetContactpersons(contactpersons []Contactperson) {
	b.Contactpersons = contactpersons
}

func (b *Businesspartner) AddContactperson(contactperson Contactperson) {
	b.Contactpersons = append(b.Contactpersons, contactperson)
}*/