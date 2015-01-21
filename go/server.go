package main

import (
	"log"

	employee "./fourth_employee"
	"code.google.com/p/goprotobuf/proto"
)

func main() {
	test := &employee.Employee{}
	test.Name = proto.String("Simon")
	test.Age = proto.Int32(37)
	test.NiNumber = proto.String("JJ817203D")
	d, _ := proto.Marshal(test)

	newTest := &employee.Employee{}

	proto.Unmarshal(d, newTest)

	log.Println(newTest)

	log.Println(d)
}
