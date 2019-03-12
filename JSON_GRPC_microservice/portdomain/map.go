package portdomain

import (
	"exercises/JSON_GRPC_microservice/jsonstream"
	"fmt"
)

type DB map[string]interface{}

var mapDB DB

func WorkDB(chanPort chan jsonstream.Port) {

	mapDB = make(map[string]interface{})

	for i := range chanPort {
		mapDB.AddUpdateKey(i)
	}
	mapDB.PrintDB()
}

func (mapdb DB) IsExistKey(key string) bool {
	if _, ok := mapdb[key]; ok {
		return true
	}
	return false
}

func (mapdb DB) AddUpdateKey(port jsonstream.Port) {
	mapdb[port.Key] = port.Message
}

func (mapdb DB) PrintDB() {
	for key, value := range mapdb {
		fmt.Println("Key:", key, "Value:", value)
	}
}
