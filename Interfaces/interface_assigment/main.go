package main

import "fmt"

type shape interface{
	getArea() float64
}

type square struct{
	sideLength float64
}

type trinagle struct{
	height float64
	base float64
}

func main (){
	s := square { sideLength: 10}
	t := trinagle { height: 10, base: 20}

	fmt.Println("Square:", s);
	fmt.Println("Triangle:", t);

	printArea(s)
	printArea(t)
	
}

func( s square) getArea() float64{
	return s.sideLength*4
}

func( t trinagle) getArea() float64{
	return (t.height*t.base)
}

func printArea (s shape) {

	fmt.Println(s.getArea())
}
