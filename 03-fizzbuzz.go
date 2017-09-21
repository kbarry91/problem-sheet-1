//  Author: Kevin Barry
//  Date:   21/09/2017
//  Adapted from: http://wiki.c2.com/?FizzBuzzTest
package main

	import "fmt"//import for input/output

	func main() {
		
		fmt.Println("starting fizzbuzz")
		c := make([]int, 100)

	for i := range c {
		d := i + 1
		threes := d%3 == 0//boolean named threes to determine if value is divisble by 3
		fives := d%5 == 0

		if threes && fives {//if divisible by 3 and 5 true
			fmt.Println("FizzBuzz")
		} else if threes {
			fmt.Println("Fizz")
		} else if fives {
			fmt.Println("Buzz")
		} else {//if not print value
			fmt.Println(d)
		}
	}//for
}//close main
