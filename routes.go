package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Tables",
		"GET",
		"/tables",
		TablesIndex,
	},
	Route{
		"Reservation",
		"GET",
		"/reservation",
		ReservationsIndex,
	},
	Route{
		"User",
		"GET",
		"/users/{user}",
		UserByName,
	},
	Route{
		"Days",
		"GET",
		"/days",
		DaysIndex,
	},
	Route{
		"DayByOffset",
		"GET",
		"/days/{offset}",
		DayByOffset,
	},
}
