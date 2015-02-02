package main

import (
	"bytes"
	"log"
	"net/http"

	people "../fourth_people"
	"code.google.com/p/goprotobuf/proto"
)

func main() {
	test := &people.Employee{}
	test.Name = proto.String("Simon")
	test.Age = proto.Int32(37)
	test.NiNumber = proto.String("MM764987F")
	d, _ := proto.Marshal(test)

	log.Println("posting")
	resp, err := http.Post("http://localhost:8080/ReceiveEmployee", "application/x-protobuf", bytes.NewBuffer(d))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(d)
		log.Println(resp)
		log.Println(bytes.NewBuffer(d))
	}
}
