package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {

	fmt.Println("Hi it is a simple app for DevOps class")
}
