package main

import (
	"fmt"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	name := githubactions.GetInput("name")
	fmt.Printf("Hello, %v! We are running on %v; Hooray!\n", name, platform())
}
