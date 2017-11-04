//Q5. Guessing game
//User must guess a randomly generated number
//After every guess the program tells the user whether their number was too high or too low. 
//The number of tries should be printed. 
//It counts only as one try if they input the same number multiple times consecutively.


//Code addapted from the following links
//http://golang-basic.blogspot.ie/2014/05/guess-number-game-in-go.html
//https://gist.github.com/marklap/e76017eea057d433c8ed81175125a826

package main

//imports
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main(){
	//Constant variables initialized (values cannot change)
	const maxGuesses = 6 
	const maxNum = 20

	var guess, unknown int 
	cnt := 0; //initializes cnt to 0
	
	//Random number generated according to the current time
	rand.Seed(time.Now().UnixNano())
	unknown = rand.Intn(maxNum)+1
	
	
	//fmt.Println(unknown) //Check with the visible random number if the requirements are met 
	for {
		if maxGuesses != cnt { //as long as the amount of guesses the user makes (cnt) does not exceed maxGuesses execute this if statement
			cnt++; //number of guesses the user makes
			
			fmt.Println("\nYOU HAVE 6 GUESSES! GO!")
			fmt.Println("Guess a number between 1 and 10: ")
			fmt.Scanln(&guess) //input number from the user
			
			//Equal/Smaller/Bigger
			if guess == unknown{
				fmt.Printf("\nYou Won! The random number is %d and you won in only %d tries. Well done!", unknown, cnt)	//if guessed number is the same as random number print this
				os.Exit(0)
			} else if guess < unknown {
				fmt.Printf("\nToo small. Try again!") //if guessed number is smaller than random number print this
			} else if guess > unknown {
				fmt.Printf("\nToo big. Try again!") //if guessed number is smaller than random number print this
			}		
		} else { //let the user know he reached the maximum number of guesses allowed and exit the loop process
			fmt.Println("The random number was %d and you reached your max amount of tries %d \n. Better luck next time!", unknown, cnt)
			os.Exit(0)
		}//if/else..if	
			
	}//for
	
}//main
