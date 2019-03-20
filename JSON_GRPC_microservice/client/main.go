package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	pb "exercises/JSON_GRPC_microservice/api/proto/v1"
	stream "exercises/JSON_GRPC_microservice/client/pkg"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	path    = "client/example.json"
)

func grpcDial() {

}
func UploadPorts(w http.ResponseWriter, r *http.Request) {
	//	params := mux.Vars(r)
	//	log.Println(params, params["path"])

	// Set up a connection to the server.

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer conn.Close()

	c := pb.NewPortDomainClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ok, err := c.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if ok == nil || err != nil {
		json.NewEncoder(w).Encode(err)
		json.NewEncoder(w).Encode(ok)
		log.Println("GRPC is not OK", ok, err)
		return
	}

	//open file as stream
	file, err := os.Open(path)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Println(err)
		return
	}
	defer file.Close()

	chanPort := make(chan stream.Port)
	go stream.StreamJSON(file, chanPort)
	for msgJSON := range chanPort {
		msgString, _ := json.Marshal(msgJSON.Message)

		//	return skey, string(s), nil
		c.AddUpdatePort(ctx, &pb.AddUpdatePortRequest{Key: msgJSON.Key, Data: msgString})
		if err != nil {
			log.Printf("error: %v", err)
			json.NewEncoder(w).Encode("done")
		}
		//log.Printf("ADD: %v %v", msgJSON, string(msgString))
	}
	json.NewEncoder(w).Encode("done")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/ports/upload", UploadPorts).Methods("GET")
	//router.HandleFunc("/upload/{path}", UploadDB).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
