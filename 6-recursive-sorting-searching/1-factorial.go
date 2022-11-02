package main

import "fmt"

/*
5! = 5*4*3*2*1 = 120
	= 5*4!

4! = 4*3*2*1
	= 4*3!
3! = 3*2!
2! = 2*1!
1! = 1
*/

func factorial(n int) int {
	if n == 1 { // base case
		return 1
	} else { //recurrence relation
		/*
			5 * f(4) --> 5*4! --> 5*24=120
			4 * f(3) --> 4*3! -->4*6=24
			3 * f(2) -->3*2=6
			2 * f(1) -->2*1=2
			1
		*/
		return n * factorial(n-1)
	}
}

func factorialLoop(n int) int {
	var hasil int = 1
	for i := n; i >= 1; i-- {
		hasil = hasil * i
	}
	return hasil
}

func main() {
	var data = 5
	fmt.Println(factorial(data))     // 120
	fmt.Println(factorialLoop(data)) // 120

}
