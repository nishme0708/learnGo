package main

import "fmt"

func main() {
	var numbers [4]int
	fmt.Printf("%v\n", numbers)
	fmt.Println(numbers)

	var a1 = [4]int{1, 2, 3, 4}
	fmt.Println(a1, len(a1))

	var cities []string
	fmt.Println("cities is equal to nil", cities == nil)
}
