/*
	Author: Kevin Barry
	Date:   22/09/2017
	Program: Determine if a string is a palindrome
*/
package main

import "fmt"	//import for input/output
import "strings"//import to work with strings
func isPalindrome(input string) bool {
	//replace upper case letters with lowercase
	input = strings.ToLower(input) 
	
	//loop through half string
	for i := 0; i < len(input)/2; i++ {
		//check if char equals the adjacent char
		if input[i] != input[len(input)-i-1] {
			return false
		}//if
	}
	return true	
}//isPalindrome

func main() {
	
	myStr := ""
	fmt.Print("Enter String to check palindrome: ")
	fmt.Scanf("%s", &myStr)

	//run isPalindrome test and print output
	if isPalindrome(myStr) == true {
		fmt.Printf("'%s' is a Palindrome",myStr)
	}else{
		fmt.Printf("'%s' is NOT a Palindrome",myStr)
	}

}//close main