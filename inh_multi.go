package main

import (
	"context"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

func resolver() {
	// Discover all services on the network (e.g. _workstation._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			log.Println(len(entries))
			log.Println(entry)
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err = resolver.Browse(ctx, "_tttt._tcp", "local.", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
}
func main() {
	meta := []string{
		"version=0.1.0",
		"hello=word",
	}
	service, _ := zeroconf.Register(
		"tttt",
		"_tttt._tcp",
		"local.",
		3333,
		meta,
		nil,
	)
	defer service.Shutdown()

}
