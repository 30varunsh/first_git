package main

import "fmt"

// func main() {
// 	colors := map[string]string{//map[key]value{}
// 		"red":   "#ff0000",
// 		"green": "#4bf745",
// 	}
// 	fmt.Println(colors)
// }

//-------

// func main() {
// 	colors := make(map[string]string)
// 	colors["white"] = "#ffffff"
// 	delete(colors, "white")//delete function in  map

// 	fmt.Println(colors)
// }

//-------

func main() {
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	color["yellow"] = "#fff000"
	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex Code of", color, "is", hex)
	}
}
