//Q2. Current time
	//Program that prints the 
	//current time and date to the console.
	
	
//Code addapted from
	//https://tour.golang.org/welcome/4

package main

//imports paths from package main
import (
	"fmt"
	"time" 	
)

func main() {
	fmt.Println("Welcome to the playground!") //prints out the string

	fmt.Println("The time is", time.Now()) //outputs current time
}//main
