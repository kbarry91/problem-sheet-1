//  Author: Kevin Barry
//  Date:   27/09/2017
//  Description : function to calculate the square root of a number using Newton’s method.
package main

	import {
import "fmt"
import "fmt"
		"fmt"//import for input/output
		"math"// import for math (math.Sqrt(x))
	}//imports

	//function to run newthons method
	func z_next(z float64,x float64)float64{
		return z - (((z*x) - x) / (2*z))
	}//z_next

	func main(){
		//number to calculate square root
		x := 20.0

		//initial guess
		z :=  float64(1)

		//loop through using newtons method until next guess == current guess
		for z = 1.0; z != z_next(z,x); z = z_next(z,x){
			//display current guess
			fmt.Println("The current Guess is %12.8f \n ",z)
		}//for

		fmt.Println("The square of %f")
	}//main