package domain

import "gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/datatype"

type Contactperson struct {
	ID uint `gorm:"primary_key"`
	BPRefer uint
	UUID string
	Title string
	Prename string
	Lastname string
	Email datatype.Email `gorm:"ForeignKey:EmailRefer"`
	EmailRefer uint
	Telephone string //vt
}

func NewContactperson(title, prename, lastname, telephone string, email datatype.Email) *Contactperson {
	return &Contactperson{Title: title, Prename: prename, Lastname: lastname, Email: email, Telephone: telephone}
}

func (c *Contactperson) SetTitle(title string) {
	c.Title = title
}

func (c *Contactperson) SetPrename(prename string) {
	c.Prename = prename
}

func (c *Contactperson) SetLastname(lastname string) {
	c.Lastname = lastname
}

func (c *Contactperson) SetTelephone(telephone string) {
	c.Telephone = telephone
}

func (c *Contactperson) SetEmail(email datatype.Email) {
	c.Email = email
}