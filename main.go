package main

import "fmt"

func main() {
	cards := newDeckFromFile("cards.txt")

	fmt.Println(cards)
}
