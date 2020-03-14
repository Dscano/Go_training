package main

import "fmt"

//Programm for evaluating if a number is even or odd
func main() {

	numbers := []int{}

	for i := 0; i < 11; i++ {
		numbers = append(numbers,i)
	}


	for j := range numbers {
		if numbers[j]%2 == 0 {
			fmt.Println(numbers[j] , "is even")
		} else {

			fmt.Println(numbers[j] , "is odd")
		}
	}
}