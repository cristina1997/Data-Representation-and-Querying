//Question 7
	//Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.
	
	
//Code addapted from the following links
	//https://github.com/jdxcode/programming-in-go/blob/master/05-procedural/palindrome.go
	//https://play.golang.org/p/5FUOzjBa-o

package main

//imports paths from package main
import (
	"fmt"
	"strings"
)

func main() {
	var str string; //string variable
	fmt.Println("Enter string to see if palindrome:")
	fmt.Scanf("%s", &str) //inputs a string str	

	if palindrome(strings.ToLower(str)) == true {
		//str is input
		//if palindrome(str) returns true then str is a palindrome
		//strinngs.ToLower(str) sends the value of the inputed string to the palindrome() func as lower cases
			//strings.ToLower - changes string characters to lower cases
			//strings.ToUpper - changes string characters to upper cases
		fmt.Printf("\n%s is a palindrome", str)

	} else {
		//if palindrome(str) returns false then str is not a palindrome
		fmt.Printf("\n%s is not a palindrome", str)	

	}//if/else..if
	
}//main

func palindrome(s string) bool {
	
	for i := 0; i < len(s)/2; i++ {
		//divides string inputed s into 2 identical sides every time and compares 
		//the 2 ends of it each and every time using the next if statement until it gets to the middle
		//similar to mergeSort

		if s[i] != s[len(s)-i-1] {			
			//if the first character in string s inputed from main() is not the same 
			//as last character from the same string return s
			return false
		}//if
	}//for
	return true
}//reverse



