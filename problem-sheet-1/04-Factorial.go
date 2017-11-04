//Q4. Factorial digit sum
//Function that calculates the sum of the digits of the factorial of a number. n! 
//means n x (n − 1) ... x 3 x 2 x 1. 
//For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, and the sum of the digits is 27. 
//Find the sum of the digits in the number 100!.


//Code addapted from the following links
	//Factorial
//http://www.golangprograms.com/go-program-to-find-factorial-of-a-number.html
//https://stackoverflow.com/questions/9302681/c-how-to-break-apart-a-multi-digit-number-into-separate-variables

	//Big Int
//https://stackoverflow.com/questions/17564335/golang-math-big-what-is-the-max-value-of-big-int
//https://stackoverflow.com/questions/11270547/go-big-int-factorial-with-recursion
		
package main

//imports
import (
	"fmt"
	"math/big"
)	

func main(){
	//Variables	
	var i int //numbers 0-100 only so there is no need for any other variable
	
	//Assign big int variables to variables above the highest int value existent (uint64)
	verybigfact := big.NewInt(1) //initializes verybigfact to 1
	sum := big.NewInt(0) //initializes sum to 0
	cnt := big.NewInt(0) //initializes sum to 0
	ten := big.NewInt(10) //initializes ten to 10
	fact := new(big.Int) //creates a new big int
	digit := new(big.Int) //creates a new big int	
	
	//Calculate Factorial for !10 and if it works try the same for !100
	/*
		for i = 10; i > 0; i-- {		
			fact *= i
					
		}//for
		
		fmt.Println(fact)
		
		//Get digits
		for fact > 0 {
			digit = fact % 10
			fact /= 10
			sum += digit
		}	
		fmt.Println(sum);
	*/
	
	//Make factorial a very big number
	for i = 0; i < 100; i++{ //if i between 0-100 keep looping (increments from 0 with every loop)
		fact = big.NewInt(1) //assigns big.NewInt to fact variable
		fact.Mul(verybigfact, ten) //multiplies a verybigfact with ten assigned to number 10 above
		verybigfact = fact //give verybigfact the value of fact
	}	
	fmt.Println(verybigfact) //output verybigfact

	//Calculate Factorial
	for i = 100; i > 0; i-- { //if i 100-0 (decrements from 100 with every loop)
		cnt.Add(cnt, big.NewInt(1)) // add cnt to one (increment by 1 each time)		
		verybigfact.Mul(verybigfact, cnt) //multiply verybigfact from above with cnt
				
	}//for
	
	for verybigfact.BitLen() > 0 { //if the length of verybigfact number > 0 execute
		digit.Mod(verybigfact, ten) //modulus of verybigfact with ten is digit
		verybigfact.Div(verybigfact, ten) //divide verybgfact to ten to give you the digits
		sum.Add(sum, digit) //add all digits together
	}//for	
	fmt.Println(sum) //output sum of digits
	
}//main
