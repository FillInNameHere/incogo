package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/database"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain/repo"
)

var ps database.Persistences
var tx database.Transactions
var bpREPO repo.BusinesspartnerREPO

func Init(p database.Persistences, t database.Transactions) *httprouter.Router {
	ps = p
	tx = t
	bpREPO = repo.NewBusinesspartnerREPO(ps)
	return Router()
}

func businesspartnerByName(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()
	name := q.Get("name")

	var bp domain.Businesspartner

	tx.Execute(func() {
		bp = bpREPO.FindByName(name)
	})

	fmt.Fprintf(w,"%s", bp)
}

func businesspartnerById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "%s", "not implemented!")
}

func saveBusinesspartner(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	bp := domain.Businesspartner{}
	decoder.Decode(&bp)
	defer r.Body.Close()
	//bp2 := domain.Businesspartner{Name:"Christian",Private:true,Addresses:[]domain.Address{domain.Address{Street:"Teststreet"}}}
	tx.Execute(func() {
		bpREPO.Save(bp)
	})

	fmt.Fprintf(w, "%s", bp)
	//fmt.Fprintf(w,"%s",r.Body)
}