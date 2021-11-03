package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"

	delete(colors, "green")

	for color, hex := range colors {
		fmt.Println(color, hex)
	}

	fmt.Println(colors)
}
