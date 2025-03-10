package main

import (
	"fmt"
	"time"
	"math"
)

func add(x int, y int) int {
	return x + y;
}

func add2(x , y int) int {
	return x + y;
}
func sub(x,y int) int {
	return x - y;
}

func swap(x,y string) (string, string) {
	return y, x;
}

func simple(x, y string) (string, string) {
	return x , y;
}

// nmaed return, here value will be automatically returned. 

func sum(x , y int) (value int) {
	value = x + y;
	return;
}

var c , python bool

func main(){

	var num int;

	anum := 55;
	
	fmt.Println("Hello Go!");
	fmt.Println("Time is ", time.Now());
	fmt.Println("Square root of the 7 is : ", math.Sqrt(7));
	fmt.Println("Lets see the pi:", math.Pi);
	fmt.Print(add(5,3));
	fmt.Print(add2(5,5));
	fmt.Print(sub(5,2));
	fmt.Println(swap("Hello", "Go"));
	
	a,b := simple("one","two"); 
	fmt.Println(a,b);
	fmt.Println(sum(10,10));

	//print the variables,
	c = true;

	fmt.Println(c, python,num, anum);
	fmt.Pritnln(type(anum));	

}


