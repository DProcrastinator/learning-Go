package main

import "fmt"

import "rsc.io/quote"


func main() {

// A slice is a data structure similar to an array, but it can change in size.
  var myslice [] string; //a slice a string
//   initialize the slice with values
  var newslice = []string {"first" , "second"}
   //or 
   greetings := []string {"hi" , "hello"}
//    an empty slice of a specific length using the make() function:
 
  test := make([]int , 3)
  
//   new slice from an existing slice, appending one or more items to it
mySlice  := []string{"First", "Second", "Third"}

newslice := append(mySlice , "fourth" , "fifth")

//duplicating using copy
mySlice := []string{"First", "Second", "Third"}

newSlice := make([]string, 3)

copy(newSlice, mySlice)

//an empty slice with capacity 10
newSlice := make([]string, 0, 10)

// get a portion of a slice
mySlice := []string{"First", "Second", "Third"}

newSlice := mySlice[:2] //get the first 2 items
newSlice2 := mySlice[2:] //ignore the first 2 items
newSlice3 := mySlice[1:3] //new slice with items in position 1-2

//maps

// A map is a very useful data type in Go.
// In other language it’s also called dictionary or hash map or associative array.
// create map
agesMap := make(map[string]int)

// add a new item to the map
agesMap["flavio"] = 39

// initialize the map with values
agesMap := map[string]int{"procrastinator": 20}

// get the value associated with a key
age := agesMap["procrastinator"]

//delete an item
delete(agesMap, "procrastinator")
}