package main

import (
	"log"
	"net/http"
)

var counter int

type h struct {
	counter int
}

func (h *h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.counter = 1
	h.counter += 1

	if h.counter == 2 {
		w.Write([]byte("Done."))
	} else {
		w.WriteHeader(500)
		log.Println("Error")
	}

}

func main() {
	http.ListenAndServe("localhost:8000", new(h))
}
