package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var directory string

	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) < 2 {
		println("There are no args!")

		fmt.Print("What is the directory?\n-> ")
		text, _ := reader.ReadString('\n')
		directory = strings.Replace(text, "\n", "", -1)
	} else {
		directory = os.Args[1]
	}

	fmt.Println("The directory should be:", directory)

	// tests:

	sendWebhook(directory)
}
