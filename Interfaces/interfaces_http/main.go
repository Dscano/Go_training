package main

import (

	"fmt"
	"net/http"
	"os"
	"io"
)

type longWriter struct{

}

func main(){

	resp, err := http.Get("http://google.com")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}


	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	//This line of code due the same tingh of the 3 line above
	io.Copy(os.Stdout, resp.Body)

	//Code for checking if custom Writer works
	lw := longWriter{}
	io.Copy(lw, resp.Body)

}


//Function that impmenet custom Writer
func(longWriter) Write (bs []byte) (int, error){
	fmt.Println(string(bs))
	return len(bs),nil
}