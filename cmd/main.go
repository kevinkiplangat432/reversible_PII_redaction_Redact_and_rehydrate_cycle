package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your prompt: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	prompt := strings.TrimSpace(input)

	fmt.Println("you typed:", prompt)


}