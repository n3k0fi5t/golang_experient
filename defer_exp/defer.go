package main

import "fmt"

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

/* the sequence of return and defer
1. assgin return value to named return value
2. invoke defer function
3. return to caller
*/
func main() {
	fmt.Printf("%d return %d\n", 1, f1()) // 1 => result++
	fmt.Printf("%d return %d\n", 2, f2()) // 5 => t is already pass to r then execute defer function to change t
	fmt.Printf("%d return %d\n", 3, f3()) // 1 => pass by value
}
