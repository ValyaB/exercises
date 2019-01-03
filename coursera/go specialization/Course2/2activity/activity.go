package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var value [4]float64
var v struct {
	acceleration, velocity, displacement, time float64
}

// GenDisplaceFn calculate something
func GenDisplaceFn(a, o_v, o_s float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		s := a*t*t/2 + o_v*t + o_s
		return s
	}
	return fn
}

func main() {

	fmt.Println("Enter each value of \nacceleration, initial velocity, initial displacement and time in the new line:")
	reader := bufio.NewScanner(os.Stdin)

	for i := 0; i < 4; i++ {
		reader.Scan()
		v, err := strconv.ParseFloat(reader.Text(), 64)

		if e := reader.Err(); e != nil || err != nil {
			fmt.Println("Please reenter value as float")
			i--
		} else {
			value[i] = v
		}
	}
	v.acceleration = value[0]
	v.velocity = value[1]
	v.displacement = value[2]
	v.time = value[3]
	fn := GenDisplaceFn(v.acceleration, v.velocity, v.displacement)
	fmt.Print("Result displacement: ")
	fmt.Println(fn(v.time))
}
