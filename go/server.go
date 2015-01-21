package main

import (
	"log"

	people "./fourth_people"
	"code.google.com/p/goprotobuf/proto"
)

func main() {
	test := &people.Employee{}
	test.Name = proto.String("Simon")
	test.Age = proto.Int32(37)
	test.NiNumber = proto.String("JJ817203D")
	d, _ := proto.Marshal(test)

	newTest := &people.Employee{}

	proto.Unmarshal(d, newTest)

	log.Println(newTest)

	log.Println(d)
}
