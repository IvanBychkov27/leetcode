/* Быстрое возведение в степень
    1, при n = 0
    a*a^(n-1), при n - не четное
   (a*a)^(n/2), при n - четное
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	a, n := 2, 62
	//a, n := 2, 10
	fmt.Println("rapid  =", rapidExponentiation(a, n))
	fmt.Println("rapid2 =", rapidExponentiation2(a, n))
	fmt.Println("pow    =", int(math.Pow(float64(a), float64(n))))
	fmt.Println("naive  =", naivePow(a, n))
}

/* быстрое возведение в степень
    1, при n = 0
    a*a^(n-1), при n - не четное
   (a*a)^(n/2), при n - четное
*/
func rapidExponentiation(a, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		return rapidExponentiation(a*a, n/2)
	}
	return a * rapidExponentiation(a, n-1)
}

/* быстрое возведение в степень
    1, при n = 0
    a*a^(n-1), при n - не четное
   (a*a)^(n/2), при n - четное
   (a*a*a)^(n/3), при n делящееся на 3 без остатка
*/
func rapidExponentiation2(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%3 == 0 {
		return rapidExponentiation(a*a*a, n/3)
	}
	if n%2 == 0 {
		return rapidExponentiation(a*a, n/2)
	}
	return a * rapidExponentiation(a, n-1)
}

func naivePow(a, n int) int {
	res := 1
	if n == 0 {
		return res
	}
	for i := 1; i < n+1; i++ {
		res *= a
	}
	return res
}
