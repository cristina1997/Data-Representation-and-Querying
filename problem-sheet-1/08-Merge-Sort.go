//Q8. Merge list and sort
	//Write a function that merges two sorted lists into a new sorted list, 
	//e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].


//Code addapted from the following links
	//http://austingwalters.com/merge-sort-in-go-golang/
	//https://gist.github.com/LordZamy/2adcb6d879fcef557d3d
	//https://stackoverflow.com/questions/25510958/go-lang-print-inputted-array
	//https://tour.golang.org/moretypes/11
	//https://github.com/ZachOrr/golang-algorithms/blob/master/sorting/merge-sort.go
		
package main 

//imports paths from package main
import 
	"fmt"	

func main(){
	arrayList1 := []int {1,4,6} //creates a new array list with values {1, 4, 6}
	arrayList2 := []int {2,3,5} //creates a new array list with values {2, 3, 5}

	fmt.Printf("\n%v", mergeSort(arrayList1, arrayList2)) //calls mergeSort function 

}//main

// Runs MergeSort algorithm on a sortedList single
func mergeSort(l1, l2 []int) []int{
	arrayLen := len(l1)+len(l2) //initializes arrayLen to the total size of the 2 arrays
	i, j := 0, 0 //initializes i and j to 0
	sortedList := make([]int, arrayLen) //it gives the new list the size of the arrayLen

	for k := 0; k < arrayLen; k++{ //as long as k is smaller than the size of the array keep the loop going

		if i > len(l1)-1 && j <= len(l2)-1 {
			//every time i is bigger than the size of 1st array list-1 and j smaller than the size of 2nd array list-1
				//(-1 is used because i starts is 0-5 while len of the array is from 1-6)
			sortedList[k] = l2[j] //make the element at k position from the sorted list equal to the element at j position from l2
			
			j++ //increment j because 1 element from it was used
		} else if j > len(l2)-1 && i <= len(l1)-1 {
			//every time j is bigger than the size of 2nd array list-1 and i smaller than the size of 1st array list-1
			sortedList[k] = l1[i] //make the element at k position from the sorted list equal to the element at i position from l1

			i++ //increment i because 1 element from it was used
		} else if l1[i] < l2[j] {
			//every time an element of l1 at position i is smaller than an element of l2 at position j
			sortedList[k] = l1[i] //make the element at k position from the sorted list equal to the element at i position from l1

			i++ //increment i because 1 element from it was used
		} else {
			//every time an element of l1 at position i is bigger than an element of l2 at position j - only statement that can be true given the others are false
			sortedList[k] = l2[j] //make the element at k position from the sorted list equal to the element at j position from l2

			j++ //increment j because 1 element from it was used

		} //if/else..if
		
	}//for	
	
	return sortedList; //returns the value of sortedList
}//mergeSort

