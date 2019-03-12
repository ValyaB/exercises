package main

import (
	"exercises/JSON_GRPC_microservice/jsonstream"
	"exercises/JSON_GRPC_microservice/portdomain"
	"log"
	"os"
)

func main() {
	paralleljob := 4
	chanPort := make(chan jsonstream.Port, paralleljob)

	file, err := os.Open("jsonstream/example.json")

	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	go jsonstream.StreamJSON(file, chanPort)
	//var db portdomain.DB
	portdomain.WorkDB(chanPort)

}
