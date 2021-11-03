package main

import "fmt"

func main() {
	// Arrays - Should have length
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	fruitArrs := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(fruitArrs)

	// Slices - Arrays without fix length
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
