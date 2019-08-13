package main

import (
	"fmt"
	"os"

	"github.com/mikloslorinczi/go-test/client"
)

func main() {
	fmt.Println("Go-test")
	res, err := client.Call("https://reqres.in/api/users")

	if err != nil {
		fmt.Println("Testing Request Error:", err)
		os.Exit(1)
	}

	fmt.Println("Testing Request Results:\n", res)
}
