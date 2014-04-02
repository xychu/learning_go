package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting+s.Punct+s.Who)
}

func main() {
	http.Handle("/string", String("I'm frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Ethan!"})
	http.ListenAndServe("localhost:4000", nil)
}
