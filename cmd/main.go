package main

import (
	"fmt"

	"github.com/kevinkiplangat432/arvis_v1/internal/greetings"
)

func main() {
	message := greetings.Hello("kevin")
	fmt.Println(message)
}