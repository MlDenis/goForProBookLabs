package main

import "fmt"

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println()
	fmt.Printf("Step one\n")
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Lenght: %d\n", cap(aSlice), len(aSlice))

	fmt.Println()
	fmt.Printf("Step two\n")
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Lenght: %d\n", cap(aSlice), len(aSlice))

	fmt.Println()
	fmt.Printf("Step three\n")
	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	fmt.Printf("Cap: %d, Lenght: %d\n", cap(aSlice), len(aSlice))
}
