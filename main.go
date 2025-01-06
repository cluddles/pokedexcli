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
		fields := cleanInput(str)
		if len(fields) > 0 {
			fmt.Printf("Your command was: %s\n", fields[0])
		}
	}
}
