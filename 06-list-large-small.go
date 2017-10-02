/*
	Author: Kevin Barry
	Date:   22/09/2017
	Program: Determine the largest and smallest int from a list
*/
package main

import "fmt"//import for input/output

func main() {
	//create list of random numbers
	numlist := []int{26, 42, 17, 10, 4, 67, 98, 31, 82, 56 }
	var small = numlist[0]
	var large = numlist[0]
	
	//for loop to iterate through numlist
	for i := 0; i <  len(numlist) ; i++ {

		//compare largest and smallest
		if numlist[i]>large{
			large = numlist[i]
		}else if numlist[i]<small{
			small = numlist[i]
		}

	}//for 

	//output largest and smallest
	fmt.Println("The largest number is ", large)
	fmt.Println("The smallest number is ", small)

}//close main