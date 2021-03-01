package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/day253/spellcheck"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	model := spellcheck.Train("big.txt")

	for {
		fmt.Print("Text: ")
		scanner.Scan()
		query := scanner.Text()
		if len(query) != 0 {
			fmt.Println("Result: ", spellcheck.Correct(query, model))
		} else {
			break
		}
	}
}
