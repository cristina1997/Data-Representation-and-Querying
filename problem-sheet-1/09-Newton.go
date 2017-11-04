//Q9. Newton’s method for square roots
	//Write a function to calculate the square root of a number using Newton’s method. 
	//Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.

		//z_next = z - ((z*z - x) / (2 * z))

	//This is repeated while the values of z_next and z are different. 
	//After each iteration the value of z is replaced with that of z_next.
	//Run a few tests to determine how close you are to the math.Sqrt value?

	//Hint: to declare and initialize a floating point value, give it floating-point syntax or use a conversion:
		//z := float64(1)
		//z := 1.0


//Code addapted from the following links
	//https://tour.golang.org/flowcontrol/8

package main

//imports paths from package main
import (
	"fmt"
	"math"
)

func main() {
	//Variables
	var x int //attempt with integer
	//var x float64; //attempt with float (does not work)

	fmt.Print("Please enter a number to calculate the square root: ")
	fmt.Scanf("%d", &x) //attempt with integer
	//fmt.Scanf("%f", &x) //attempt with float

	//fmt.Printf("The square root of %d is %v", x, newton(float64(x))); //float64(x) changes x from int to a float64
	fmt.Printf("The square root of %.2f is %.2f", x, newton(x)) //float
	
}//main

func newton(x float64) float64 {
	var z_next float64 //it needs to be a float because the devision can give out a result with decimal numbers
	z := float64(1) //z needs to take the value of z_next which is a float so z needs to be a float too

	for z != z_next { 
		z_next = z - ((z*z - x) / (2 * z)) //Newton's formula
		z = z_next //makes z equal to the value of z_next calculated from above

		if z == z_next {
			//the if statement keeps the for loop going until z = z_next at the end of the if loop
			//once that happens the foor loop is exited because its statement is false
			z_next = z - ((z*z - x) / (2 * z)) //if z and z_next are both equal execute the same formula again
			
		}//if		
		
	}//for	
	return z //it returns the end result of z

}//newton
