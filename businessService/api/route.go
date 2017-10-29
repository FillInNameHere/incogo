package api

import "github.com/julienschmidt/httprouter"

func Router() *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/businesspartner", businesspartnerByName)
	mux.GET("/businesspartner/:id", businesspartnerById)
	mux.POST("/businesspartner", saveBusinesspartner)
	return mux
}