package repo

import (
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/database"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain"
)

type AddressREPO struct {
	ps database.Persistences
}

func NewAddressREPO(p database.Persistences) AddressREPO {
	return AddressREPO{p}
}

func (a AddressREPO) Save(address domain.Address) {
	a.ps.Save(address)
}

func (a AddressREPO) FindById(id int) domain.Address {
	address := domain.Address{}
	a.ps.Find(id, &address)
	return address
}