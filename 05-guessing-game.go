//  Author: Kevin Barry
//  Date:   21/09/2017
//  Guessing game
package main

import "fmt"
import "math/rand"//import for random number

func main() {

	var userinput int//variable to take user input
	var randomNum int = rand.Intn(100)//generate random number
	var prevNum int
	var counter int 
	var numFound bool =false

	fmt.Println()
	fmt.Println("Number Guessing Game")
	fmt.Println("====================")

	//loop while numfound = false
	for numFound != true{
		
		//take user input from console
		fmt.Print("Please enter a number(1-100):" )
		fmt.Scan(&userinput)
		counter++

		if userinput==prevNum{
			counter--
			fmt.Println("Number already chosen")			
		}

		//compare user input to random number
		if userinput >randomNum{
			fmt.Println("Number was too large")
		}else if userinput < randomNum{
			fmt.Println("Number was too small")
		}else if  userinput == randomNum{
			fmt.Println()
			fmt.Println("Guess correct Random number was ",randomNum)
			fmt.Println("Took ",counter," attempts")
			numFound = true ;//set bool numFound to true
		}//if

		prevNum = userinput
		
	}//for
}//main