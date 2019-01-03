package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Print(c.food)
}
func (c Cow) Move() {
	fmt.Print(c.locomotion)
}
func (c Cow) Speak() {
	fmt.Print(c.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Print(b.food)
}
func (b Bird) Move() {
	fmt.Print(b.locomotion)
}
func (b Bird) Speak() {
	fmt.Print(b.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Print(s.food)
}
func (s Snake) Move() {
	fmt.Print(s.locomotion)
}
func (s Snake) Speak() {
	fmt.Print(s.noise)
}
func standardizeSpaces(s string) []string {
	s = strings.ToLower(s)
	return strings.Fields(strings.Join(strings.Fields(s), " "))

}
func main() {
	names := make(map[string]Animal)

	c1 := Cow{food: "grass", locomotion: "walk", noise: "moo"}

	b1 := Bird{food: "worms", locomotion: "fly", noise: "peep"}

	s1 := Snake{food: "mice", locomotion: "slither", noise: "hsss"}

	fmt.Print("Enter request or CTRL^C for exit\n<newanimal> <name> <type> \n<query>     <name> <info_type> \n \n>")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		request := standardizeSpaces(scanner.Text())
		if len(request) == 3 {
			switch request[0] {
			case "newanimal":
				switch request[2] {
				case "cow":
					names[request[1]] = c1
					fmt.Print("Created it!")
				case "bird":
					names[request[1]] = b1
					fmt.Print("Created it!")
				case "snake":
					names[request[1]] = s1
					fmt.Print("Created it!")
				default:
					fmt.Print("wrong type of animal")
				}

			case "query":
				_, ok := names[request[1]]
				if !ok {
					fmt.Print("name is not found\n>")
					continue
				}

				switch request[2] {
				case "eat":
					names[request[1]].Eat()
				case "move":
					names[request[1]].Move()
				case "speak":
					names[request[1]].Speak()
				default:
					fmt.Print("wrong type of animal")
				}

			default:
				fmt.Print("Please type newanimal or query")
			}

		} else {
			fmt.Print("Format of request: <newanimal> <name> <type> or \n    <query> <name> <info_type>")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		fmt.Print("\n>")
	}
}
