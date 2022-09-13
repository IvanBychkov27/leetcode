/* быстрое возведение в степень
    1, при n = 0
    a*a^(n-1), при n - не четное
   (a*a)^(n/2), при n - четное
*/

package main

import "fmt"

func main() {
	fmt.Println("rapid =", rapidExponentiation(2, 10))
}

/* быстрое возведение в степень
    1, при n = 0
    a*a^(n-1), при n - не четное
   (a*a)^(n/2), при n - четное
*/
func rapidExponentiation(a, n int) int {
	res := 1
	if n == 0 {
		return res
	}

	if n%2 == 0 {
		res = rapidExponentiation(a*a, n/2)
	} else {
		res = a * rapidExponentiation(a, n-1)
	}
	return res
}
