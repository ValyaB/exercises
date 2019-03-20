package stream

import (
	"encoding/json"
	"io"
	"log"
)

type Port struct {
	Key     string
	Message interface{}
}

func StreamJSON(file io.Reader, chanPort chan Port) {

	defer close(chanPort)

	stream := json.NewDecoder(file)
	// read open bracket
	t, err := stream.Token()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("INFO: %T: %v\n", t, t)

	for stream.More() {

		var skey string
		var message interface{}

		key, err := stream.Token()

		if err != nil {
			log.Println(err)
			continue
		}

		switch value := key.(type) {
		case string:
			skey = value
		default:
			log.Println("WARNING: not string", value)
			continue
		}

		if err := stream.Decode(&message); err == io.EOF {
			break
		} else if err != nil {

			log.Println(err)
			continue
		}
		//		log.Printf("%v\n", message)
		if skey == "" {
			log.Println("INFO: drop Port with empty key string", skey)
			continue
		}
		chanPort <- Port{skey, &message}
	}
	// read closing bracket
	t, err = stream.Token()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("INFO: %T: %v\n", t, t)

}
