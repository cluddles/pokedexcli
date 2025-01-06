package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		runCommand(str)
	}
}
