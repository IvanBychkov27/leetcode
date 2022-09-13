/*
   Алгоритм Евклида - нахождение НОД (наибольшего общего делителя)
*/
package main

import "fmt"

func main() {
	a := 1250
	b := 125
	res := gcd(a, b)
	fmt.Println("gcd =", res)
}

// нахождение наибольшего общего делителя алгоритмом Евклида
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
