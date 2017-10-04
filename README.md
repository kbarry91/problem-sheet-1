## problem-sheet-1

This repository was constructed to upload my programs for "problem sheet 1" based on the Go programming language.

### To run the Programs
Install Go on your device by visiting https://golang.org/

Clone the repository in cmd by executing the following command 

```git clone https://github.com/kbarry91/problem-sheet-1.git```

Now you can build an .exe file with the command
```go build <filename.go> ```
and run the program with
```<filename>``` 
or simply use
```go run <filename>```

### Problem set: Go fundamentals
#### 1. Kon’nichiwa, Sekai!
Write a program that prints Hello, world! in Japanese (using Japanese characters) to the screen.

#### 2. Current time
Write a program that prints the current time and date to the console.

#### 3. FizzBuzz
Write a program that prints the numbers from 1 to 100, each on a new line, to the console, except for the following conditions. For multiples of three print Fizz instead of the number, and for multiples of five print Buzz. For numbers that are multiples of both three and five print FizzBuzz.

#### 4. Factorial digit sum
Write a function that calculates the sum of the digits of the factorial of a number. n! means n x (n − 1) ... x 3 x 2 x 1. For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. Then find the sum of the digits in the number 100!.

#### 5. Guessing game
Write a guessing game where the user must guess a randomly generated number. After every guess the program tells the user whether their number was too high or too low. At the end, the number of tries should be printed. It counts only as one try if they input the same number multiple times consecutively.

#### 6. Largest and smallest in list
Write a function that returns the largest and smallest elements in a list.

#### 7. Palindrome test
Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.

#### 8. Merge list and sort
Write a function that merges two sorted lists into a new sorted list, e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].

#### 9. Newton’s method for square roots
Write a function to calculate the square root of a number using Newton’s method. Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.

z_next = z - ((z*z - x) / (2 * z))

#### 10. Reverse string
Write a function to reverse a string in Go.
