/*	Author: Kevin Barry
  	Date:   02/10/2017
  	Reverse String
*/

package main

import "fmt"

//function to reverse string
func revStr(myStr string) string {
	revString := ""
	// Loop through the string backwards and store in a string
    for i := len(myStr)-1; i >= 0; i-- {
        revString += string(myStr[i])
    }
    return revString 
}

func main() {
	myStr := "This is a string to be reversed" 

	//print output
	fmt.Printf("Original String : %s \n", myStr) 
	fmt.Printf("Reversed String : %s",revStr(myStr))
}//main