package routes

import (
	"fmt"
	"net/http"
)

func HandleRequests() {
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})
}
