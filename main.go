package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tables", TablesIndex)
	router.HandleFunc("/reservations", ReservationsIndex)
	router.HandleFunc("/users/{user}", UserByName)
	router.HandleFunc("/days", DaysIndex)
	router.HandleFunc("/days/{offset}", DayByOffset)

	log.Fatal(http.ListenAndServe(":8080", router))
}
