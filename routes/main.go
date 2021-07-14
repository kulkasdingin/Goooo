package routes

import "github.com/gorilla/mux"

var Router *mux.Router

const ra = "/api" // Routes Api

func Main() {
	HandleApiRequests()
	HandleRequests()
}
