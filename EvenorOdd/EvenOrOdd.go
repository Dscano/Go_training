package main

import "fmt"

//Programm for evaluating if a number is even or odd
func main() {

	numbers := []int{}

	for i := 0; i < 11; i++ {
		numbers = append(numbers,i)
	}


	for _, number := range numbers {
		if number %2 == 0 {
			fmt.Println(number , "is even")
		} else {

			fmt.Println(number , "is odd")
		}
	}
}