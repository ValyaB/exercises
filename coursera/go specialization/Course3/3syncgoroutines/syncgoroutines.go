package main

import (
	"fmt"
	"sync"
)

func host(req chan bool) {
	for {
		req <- true
	}
}
func philosoferEat(number int, philo philosofer, c chan bool) {
	for i := 0; i < 3; i++ {
		<-c
		//Go’s map iteration order (using the range keyword) is random
		//https://blog.golang.org/go-maps-in-action
		for i, v := range philo {
			v.Lock()
			fmt.Println("lock", i)
		}
		fmt.Println("starting to eat N", number)
		c
		for i, v := range philo {
			v.Unlock()
			fmt.Println("lock", i)
		}
	}
	fmt.Println(philo)
}

type philosofer map[string]*sync.Mutex

func main() {
	req := make(chan bool, 2)
	host(req)
	chops := [5]sync.Mutex{}
	var philos [5]philosofer

	for i := range philos {
		philos[i] = philosofer{"leftCS": &chops[i], "rightCS": &chops[(i+1)%5]}
	}
	for i := range philos {
		philosoferEat(i+1, philos[i], req)
	}
	fmt.Println(philos)

}
