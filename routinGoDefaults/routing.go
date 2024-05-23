package routingodefaults

import (
	"fmt"
	"net/http"
)

func RoutingGo() {
	fmt.Println("Building rest api in go 1.22")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Return all comments")

	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Return a single comment: %s", id)

	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "POST A NEW COMMENT")

	})

	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
