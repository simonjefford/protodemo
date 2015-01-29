package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	people "./fourth_people"
	"code.google.com/p/goprotobuf/proto"
	"github.com/gorilla/mux"
)

func main() {
	startHttpServer()
}

func sendError(w http.ResponseWriter, v interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, v)
	log.Println(v)
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength <= 0 {
		sendError(w, "no bytes received")
		return
	}

	b := make([]byte, r.ContentLength)
	_, err := r.Body.Read(b)

	if err != nil && err != io.EOF {
		sendError(w, err)
		return
	}

	e := &people.Employee{}
	err = proto.Unmarshal(b, e)

	if err != nil {
		sendError(w, err)
		return
	}

	log.Println("Parsed output:", e)

	w.WriteHeader(http.StatusOK)
}

func startHttpServer() {
	r := mux.NewRouter()
	r.HandleFunc("/ReceiveEmployee", employeeHandler)
	log.Println("starting")
	log.Fatal(http.ListenAndServe(":8080", r))
}
