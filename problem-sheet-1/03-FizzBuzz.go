//Q3. FizzBuzz
	//Program printing numbers from 1 to 100, each on a new line, to the console, except for the following conditions. 
	//For multiples of three print Fizz instead of the number
	//For multiples of five print Buzz. 
	//For numbers that are multiples of both three and five print FizzBuzz.

	
//Code addapted from
	//http://wiki.c2.com/?FizzBuzzTest

package main

//imports paths from package main
import (
	"fmt"
) 

func main() {
	//Variable
	var output string //string used for string variabes
	
	for i := 0; i <= 100; i++{ //as long as i is between 0 and 100 keep looping (increments from 0 with every loop)
		
		if i%3==0 {
			output = "Fizz"
			fmt.Println(output) //output Fizz if i is divisible by 3
		} else if i%5==0{
			output = "Buzz"
			fmt.Println(output) //output Buzz if i is divisible by 5
		} else if i%3==0 && i%5==0 {
			output = "FizzBuzz" 
			fmt.Println(output) //output FizzBuzz if i is divisible by both 3 AND 5
		} else{
			fmt.Println(i) //if none of the other 3 conditions from above are true just output the number
		}//if/else..if
	
	}//for
	
}//main
