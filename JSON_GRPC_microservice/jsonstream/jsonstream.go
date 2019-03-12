package jsonstream

import (
	"encoding/json"
	"io"
	"log"
)

type Port struct {
	Key     string
	Message *Message
}
type Message struct {
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    string        `json:"province"`
	Timezone    string        `json:"timezone"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

func StreamJSON(file io.Reader, chanPorts chan Port) {

	defer close(chanPorts)

	stream := json.NewDecoder(file)
	// read open bracket
	t, err := stream.Token()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("INFO: %T: %v\n", t, t)

	for stream.More() {
		var message Message
		var skey string

		key, err := stream.Token()

		if err != nil {
			log.Panicln(err)
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

		chanPorts <- Port{skey, &message}

	}
	// read closing bracket
	t, err = stream.Token()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("INFO: %T: %v\n", t, t)

}
