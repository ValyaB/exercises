package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	p1 := map[string]string{"name": "", "address": ""}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name:")
	scanner.Scan()
	p1["name"] = scanner.Text()
	fmt.Println("Enter your address:")
	scanner.Scan()
	p1["address"] = scanner.Text()
	jsn, err := json.Marshal(p1)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}

	fmt.Println("JSON object is:")
	fmt.Println(string(jsn))
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
