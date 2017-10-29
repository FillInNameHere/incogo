package repo

import (
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/database"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain"
)

type ContactpersonREPO struct {
	ps database.Persistences
}

func NewContactpersonREPO(p database.Persistences) ContactpersonREPO {
	return ContactpersonREPO{ps:p}
}

func (c ContactpersonREPO) Save(cp domain.Contactperson) {
	c.ps.Save(&cp)
}

func (c ContactpersonREPO) FindByID(id int) domain.Contactperson {
	cp := domain.Contactperson{}
	c.ps.Find(id, &cp)
	return cp
}