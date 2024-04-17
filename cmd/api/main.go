package main

import (
	"fmt" //format package
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Ishanibhatia/GO-API-check-coins/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true) //when we print something out we get the file and their line number 

	var r *chi.Mux = chi.NewRouter() //we created a new chi mux variable using the new router function which returns a pointer to a mux type which is basically a struct which we'll use to set up our API

	handlers.Handler(r)

	//these are some endpoint statements
	fmt.Println("Starting GO API service...")

	fmt.Println(`GO API`)

	err := http.ListenAndServe("localhost:8000",r) //this takes the base address of our server which in our case is localhost:8000 and a handler which are muxtype 
	if err != nil{
		log.Error(err)
	}
}