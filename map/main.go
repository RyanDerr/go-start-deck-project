package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#1234ff",
	}
	addColor(colors, "white", "#ffffff")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func addColor(m map[string]string, k string, c string) {
	m[k] = c
}
