package mapdb

import (
	"encoding/json"
	"log"
)

type DB map[string]interface{}

var mapDB DB

type Port struct {
	Key     string
	Message interface{}
}

func (mapdb DB) isExistKey(key string) bool {
	if len(mapDB) == 0 {
		return false
	}
	if _, ok := mapdb[key]; ok {
		return true
	}
	return false
}

func (mapdb DB) addUpdateKey(p interface{}) (ok bool) {

	if pkey, ok := p.(Port); ok {
		mapdb[pkey.Key] = pkey.Message
	}
	return ok
}
func (mapdb DB) printPort(s string) interface{} {

	if len(mapDB) == 0 {
		return nil
	}
	if port, ok := mapdb[s]; ok {
		return port
	}
	return nil
}

func ConnectToDB() bool {
	if mapDB == nil {
		mapDB = make(map[string]interface{})
	}
	return true
}

func IsExistPort(s string) bool {
	return mapDB.isExistKey(s)
}

func GetPort(s string) interface{} {
	return mapDB.printPort(s)
}
func AddUpdatePort(key string, msg []byte) bool {
	var message interface{}

	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(err)
	}

	return mapDB.addUpdateKey(Port{Key: key, Message: message})
}

func IsAlive() bool {
	if len(mapDB) == 0 {
		return false
	}
	return true
}
