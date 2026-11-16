package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter a prompt: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error capturing input")
		return
	}

	prompt := strings.TrimSpace(input)

	fmt.Println(prompt)
}