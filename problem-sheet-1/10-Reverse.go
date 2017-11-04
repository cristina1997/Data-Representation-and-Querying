//Question 10
//Write a function to reverse a string in Go.


//Code addapted from the following links
//http://golangcookbook.com/chapters/strings/reverse/


package main

//imports paths from package main
import (
	"fmt"
	"strings"
)

func main() {
	var str string //string variable
	fmt.Println("Enter string to reverse:")
	fmt.Scanf("%s", &str) //input a string	

	if strings.EqualFold(str, reverse(str)) == true { 
		//strings.EqualsFold is used to ignore upper and lower cases of a string 
		//strings.EqualsFold takes 2 parameters and returns boolean (true or false) - best used for if statements	
				
		fmt.Printf("\nThe reverse of %s is %v", str, reverse(str)) //output the string and its reverse
	} 
}//main

func reverse(s string) string {
	char := []rune(s) //iterates through the array s

	for i, j := 0, len(char)-1; i < j; i, j = i+1, j-1 { //i initialized to 0 and j initialized to the size of char array-1 and it loops until i and j both meet

		char[i], char[j] = char[j], char[i] //initializes char[i] to char[j] and char[j] to char[i] all at the same time
	}//for

	return string(char) //returns char as a string
}//reverse
