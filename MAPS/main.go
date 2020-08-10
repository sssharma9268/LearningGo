package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#57bf86",
	}

	//var colors map[string]string
	//colors := make(map[string]string)
	colors["white"] = "#ffffff"
	//delete(colors, "white")
	printMap(colors)
	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, "-->", value)
	}
}
