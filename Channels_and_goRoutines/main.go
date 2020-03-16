package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	//Slice of links to check
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _,link := range links {
		// go means run this function inside a new routine, 
		// e.i, thread go routine
		go checkLink(link, c)
	}

	/*
	for i :=0; i< len(links); i++{
		fmt.Println(<-c)
	}
	//continuously checking 
	for {
		go checkLink(<-c,c)
	}
	*/

	//continuously checking with certain pause time
	for {
		//function literal
		go func(link string){
			time.Sleep(5 * time.Second) // pause the routine 
			checkLink(link, c)
		}(<-c)
		
	}
		
}

//Function for checking links
func checkLink(link string, c chan string){
	_, err := http.Get(link)

	if err != nil{
		fmt.Println(link, "migth be down!")
		c <-link
		return
	}

	fmt.Println(link, "is up!")
	c <-link

}