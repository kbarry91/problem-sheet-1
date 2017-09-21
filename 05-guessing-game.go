//  Author: Kevin Barry
//  Date:   21/09/2017
//  Guessing game
package main

import "fmt"
import "math/rand"//import for random number

func main() {

	var userinput int//variable to take user input
	var randomNum int = rand.Intn(100)//generate random number
	var counter int 
	var numFound bool =false
	var numEntered bool = false
	var numsPicked [50]int
	var i int = 0

	fmt.Println()
	fmt.Println("Number Guessing Game")
	fmt.Println("====================")

	for numFound != true{//run while numfound = false
		
		//take user input 
		fmt.Print("Please enter a number(1-100):" )
		fmt.Scan(&userinput)
		//counter = counter +1//increment counter
		fmt.Println(counter)
		i++
		numsPicked[i] = userinput

		for numEntered!=true{
			for y:=0 ;y< 50; y++{
				
			
				if numsPicked[y]==userinput{
					fmt.Println(counter)
					counter++
					numEntered=true
				}
			}
		}
		/*
		for y:=0 ;y< 50; y++{
			
		
			if numsPicked[y]==userinput{
				fmt.Println(counter)
				//counter--
			}
			
		}
		*/
		
		if userinput >randomNum{
			fmt.Println("Number was too large")
		}else if userinput < randomNum{
			fmt.Println("Number was too small")
		}//else if

		if userinput == randomNum{
			fmt.Println()
			fmt.Println("Guess correct Random number was ",randomNum)
			fmt.Println("Took ",counter," attempts")
			
			numFound = true ;//set bool numFound to true
		}//if
			numFound=false
	}//for
	
}