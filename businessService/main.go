package main

import (
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/api"
	"net/http"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/database"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain"
	"gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/datatype"
)

func main() {
	//Start DB Session
	ps, tx := database.Init(&domain.Businesspartner{}, &domain.Contactperson{}, &domain.Address{}, &datatype.Email{})

	//Start API
	mux := api.Init(ps, tx)
	server := http.Server{
		Addr:"127.0.0.1:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}