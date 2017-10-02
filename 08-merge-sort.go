//  Author: Kevin Barry
//  Date:   27/09/2017
//  Description : function merge and sort 2 lists

package main

import (
	"fmt"	//import for input/output
	"sort"	// import for sort
)

func mergeSort() {
	
		a := []int{10, 40, 60}
		b := []int{20, 30, 50}
	
	
		c := append(a, b...)
	
		
	
		sort.Ints(c)
		fmt.Println("List A: ",a)
		fmt.Println("List B: ",b)
		fmt.Println("List AB Merged and Sorted: ",c)
}

func main(){
	mergeSort()
}//main