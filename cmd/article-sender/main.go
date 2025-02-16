package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var directory string

	// dotenv
	makeEnv()

	reader := bufio.NewReader(os.Stdin)

	// get directory
	if len(os.Args) < 2 {
		// test for enviroment variables
		if env.dir != "" {
			directory = env.dir
		} else {
			fmt.Print("What is the directory?\n-> ")
			text, _ := reader.ReadString('\n')
			directory = strings.Replace(text, "\n", "", -1)
		}
	} else {
		directory = os.Args[1]
	}

	fmt.Println("The directory should be:", directory)

	printDirContents(directory)
	// sendWebhook(directory)
}
