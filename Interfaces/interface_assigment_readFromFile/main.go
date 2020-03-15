package main

import (

	"fmt"
	"os"
	"io"
)

// Program that alloy to read the conctent of a certain file
func main(){
	file, err := os.Open(os.Args[1]) 
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)

	}

	io.Copy(os.Stdout, file)
}