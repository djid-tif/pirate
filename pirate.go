package main

import (
	"bufio"
	"fmt"
	"os"
)

type pirate struct {
	pirate string
	bateau string
	equipage int
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}