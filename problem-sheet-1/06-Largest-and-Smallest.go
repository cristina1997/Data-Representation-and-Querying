//Question 6. Largest and smallest in list
//Write a function that returns the largest and smallest elements in a list.


//Code addapted from the following links
//Range
//https://gobyexample.com/range	

//Slice
//https://blog.golang.org/go-slices-usage-and-internals
//https://blog.golang.org/slices
//https://tour.golang.org/moretypes/15

//Input Array
//https://stackoverflow.com/questions/25510958/go-lang-print-inputted-array

//List - These linkg were NOT USED in the program but they are good source links if ever needed for future linked lists
//https://stackoverflow.com/questions/28714641/what-is-the-equivalent-for-linkedlistt-in-golang
//https://golang.org/pkg/container/list/#New

package main 

//imports paths from package main
import (
	"fmt"
)

func main(){
	var small, big float64 = 0, 0 //a float64 was used ibecause it supports negative numbers, decimal point numbers as well as integers (in case the user decides to use either of them)
	var maxList int	//the size of the list can only be an integer

	fmt.Print("Please enter the size of the list: ") 
	fmt.Scanf("%d", &maxList); //input the size of the list allowed

	numArray := make([]float64, maxList) //make() allocates a zero to numArray and returns a slice that refers to that array 
										//it also gives it the length maxList

	for i := 0; i < maxList; i++{ //as long as i is smaller than the size of the numArray list keep looping
		fmt.Printf("Please enter number %d: ", (i+1))
		fmt.Scan(&numArray[i]) //input a number from the user and add it to he array
	}//for
	
	for _, num := range numArray { //allocate every value of the iterated numArray (range iterates over a slice/map) to num
		if num < small {
			//as long as num is smaller than the value of small
			small = num //give small the value of num
		} else if num > big {
			//as long as num is bigger than the value of big
			big = num //give big the value of num
		}//if/else..if

	}//for
	fmt.Printf("Smallest number is %.2f and largest number is %.2f.", small, big) //output the smallest and biggest numbers from the list

}//main		
