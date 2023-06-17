package main

import "fmt"

func main() {
	var age int = 16
	var name = "dan"
	fmt.Println("Age", age)
	fmt.Println("Name", name)

	const i int = 10
	const j = 20
	// var l float64 = 10
	var k float64 = (2.6 * j)
	fmt.Println("i,j,k", i, j, k)

	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", numbers)

	var cities = []string{"one", "two"}
	fmt.Printf("%T\n", cities)
}
