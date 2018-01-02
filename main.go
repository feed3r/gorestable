package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

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

func TablesIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Endpoint to query today's table state")
}

func ReservationsIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Endpoint to submit a table reservation")
}

func UserByName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Endpoint to query all the table reservations for user %s\n", vars["user"])
}


func DaysIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Endpoint to query all the table reservations by day")
}

func DayByOffset(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    offset, _ := strconv.Atoi(vars["offset"]) // TODO handle error properly
    fmt.Fprintf(w, "Endpoint to query all the table reservations for T+%d\n", offset)
}
