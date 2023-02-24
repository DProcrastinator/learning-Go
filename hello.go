package main

import "fmt"

import "rsc.io/quote"


func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	
	
	// we can also set value as follows
var y = 5
var msg = "Remote host found."

var foo, bar int = 100, 200

//shorthand syntax 

vehicle := "car" 
age := 52 

// Bool true or false
var is_job_failed = false

	syntax 
const var type = value
define constant  pi with as float32 with 3.14159 value
const pi float32 = 3.14159
// create constant string variable 
const error_msg string = "Docker is not installed"
fmt.Println(pi)
fmt.Println(error_msg)

//simple loop 

m := 1 
for m < 10 {
	fmt.Println("welcome %d times " , m )
	if m == 5{
		break
	}
	m ++

}
//array with user defined function
}

func arrayExample(  x string , y string )  {
	// define array named domains as string type
    var domains [2]string

	domains[0] = x 
	domains[1] = y

	i := 0
	for i < len(domains){
		fmt.Println(domains[i])
		i++;
	}

}