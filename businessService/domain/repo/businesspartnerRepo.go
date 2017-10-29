package repo

import (
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/database"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain"
)

type BusinesspartnerREPO struct {
	 ps database.Persistences
}

func NewBusinesspartnerREPO(p database.Persistences) BusinesspartnerREPO {
	return BusinesspartnerREPO{p}
}

func (b BusinesspartnerREPO) Save(bp domain.Businesspartner) {
	b.ps.Save(&bp)
}

func (b BusinesspartnerREPO) FindByName(name string) domain.Businesspartner {
	bp := domain.Businesspartner{}
	b.ps.FirstWhere(&bp, "name = ?", name)
	return bp
}