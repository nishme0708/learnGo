package main

import "fmt"

func paintCost(height float64, width float64, rate float64) (float64, error) {
	if height <= 0 {
		return 0, fmt.Errorf("height cannot be zero or neg")
	}
	if width <= 0 {
		return 0, fmt.Errorf("width cannot be zero or neg")
	}
	if rate <= 0 {
		return 0, fmt.Errorf("rate cannot be zero or neg")
	}
	return (height * width * rate), nil
}

func main() {
	total, err := paintCost(0, 10, 50)
	if err == nil {
		fmt.Print("Total cost is", total)
	} else {
		fmt.Print("Error ", err)
	}
}
