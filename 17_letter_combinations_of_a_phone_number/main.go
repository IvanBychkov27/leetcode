/* https://leetcode.com/problems/letter-combinations-of-a-phone-number/
Буквенные комбинации телефонного номера
Учитывая строку, содержащую цифры от 2 до 9 включительно, верните все возможные комбинации букв, которые может представлять число. Возвращайте ответ в любом порядке.
Ниже приведено сопоставление цифр с буквами (точно так же, как на телефонных кнопках). Обратите внимание, что 1 не сопоставляется ни с какими буквами.

Выполнено:
Время выполнения: 0 мс, быстрее, чем 100,00 % отправленных онлайн сообщений для комбинаций букв телефонного номера.
Использование памяти: 2,3 МБ, менее 11,97 % отправленных онлайн сообщений для комбинаций букв телефонного номера.

*/
package main

import "fmt"

func main() {
	//digits := "" // []
	//digits := "11" // []
	//digits := "13" // ["d","e","f"]
	//digits := "31" // ["d","e","f"]
	//digits := "23" // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	//digits := "234" //  [adg adh adi aeg aeh aei afg afh afi bdg bdh bdi beg beh bei bfg bfh bfi cdg cdh cdi ceg ceh cei cfg cfh cfi]
	//digits := "2345" // [adgj adgk adgl adhj adhk adhl adij adik adil aegj aegk aegl aehj aehk aehl aeij aeik aeil afgj afgk afgl afhj afhk afhl afij afik afil bdgj bdgk bdgl bdhj bdhk bdhl bdij bdik bdil begj begk begl behj behk behl beij beik beil bfgj bfgk bfgl bfhj bfhk bfhl bfij bfik bfil cdgj cdgk cdgl cdhj cdhk cdhl cdij cdik cdil cegj cegk cegl cehj cehk cehl ceij ceik ceil cfgj cfgk cfgl cfhj cfhk cfhl cfij cfik cfil]	//digits := "2" // ["a","b","c"]
	digits := "9"

	res := letterCombinations(digits)
	fmt.Println("res =", res)
}

func letterCombinations(digits string) []string {
	data := map[byte][]byte{
		'1': {},
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	res := make([]string, 0)

	dig := make([]byte, 0, len(digits))
	for _, d := range []byte(digits) {
		if d == '1' {
			continue
		}
		dig = append(dig, d)
	}

	l := len(dig)
	if l == 0 || l > 4 {
		return res
	}

	switch l {
	case 1:
		for _, s := range data[dig[0]] {
			res = append(res, string(s))
		}
	case 2:
		for _, s0 := range data[dig[0]] {
			for _, s1 := range data[dig[1]] {
				res = append(res, string(s0)+string(s1))
			}
		}
	case 3:
		for _, s0 := range data[dig[0]] {
			for _, s1 := range data[dig[1]] {
				for _, s2 := range data[dig[2]] {
					res = append(res, string(s0)+string(s1)+string(s2))
				}
			}
		}
	case 4:
		for _, s0 := range data[dig[0]] {
			for _, s1 := range data[dig[1]] {
				for _, s2 := range data[dig[2]] {
					for _, s3 := range data[dig[3]] {
						res = append(res, string(s0)+string(s1)+string(s2)+string(s3))
					}
				}
			}
		}
	}

	return res
}
