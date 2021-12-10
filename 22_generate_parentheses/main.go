/* https://leetcode.com/problems/generate-parentheses/

Учитывая n пар круглых скобок, напишите функцию для генерации всех комбинаций правильно сформированных круглых скобок.

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Input: n = 1
Output: ["()"]

Input: n = 2
Output: ["(())","()()"]

Выполнено:
Время выполнения: 12 мс, быстрее, чем 5,58 % отправленных онлайн-сообщений для создания скобок.
Использование памяти: 3,6 МБ, менее 18,65 % отправленных онлайн-сообщений для создания скобок.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	n := 6
	them := time.Now()
	res := generateParenthesis(n)
	now := time.Now()
	fmt.Println("res =", res)
	fmt.Println()
	fmt.Println("time:", now.Sub(them))
}

var resG []string

func generateParenthesis(n int) []string {
	resG = make([]string, 0)
	s := make([]byte, n*2)
	genStr(s, 0, n*2)

	return resG
}

// генерация всех комбинаций
func genStr(s []byte, k, n int) {
	if k > n-1 {
		if check(s) {
			resG = append(resG, string(s))
		}
	} else {
		s[k] = '('
		genStr(s, k+1, n)
		s[k] = ')'
		genStr(s, k+1, n)
	}
}

// проверка комбинации на корректность
func check(data []byte) bool {
	sim := map[byte]byte{
		'(': ')',
		')': '(',
	}

	f := false
	pred := make([]byte, 0, len(data))
	for i, d := range data {
		if d == '(' {
			pred = append(pred, d)
			continue
		}
		if len(pred) == 0 || pred[len(pred)-1] != sim[d] {
			break
		}
		pred = pred[:len(pred)-1]
		if i == len(data)-1 && len(pred) == 0 {
			f = true
		}
	}

	return f
}

/* Алгоритм
Перебор всех комбинаций из произвольного кол-ва эл-тов

{
	Создать массив S[1..n];
	S := [NULL, NULL, ..., NULL];
	ProcessGeneralizedStrings (S, 1, n);
}

ProcessGeneralizedStrings(S, k, n)
{
	if k > n then Process(S)
	else
	{
	 	for i:=0 to m-1 do
	 	S[k] := ai;
	 	ProcessGeneralizedStrings (S, k-1, n);
	}
}

*/

//=================================================

func generateParenthesis_1(n int) []string {
	data := make([]byte, 0, 1024)

	for i := 0; i < n; i++ {
		data = append(data, '(', ')')
	}

	if len(data) == 0 {
		return []string{}
	}

	res := make([]string, 0, 1024)

	sim := map[byte]byte{
		'(': ')',
		')': '(',
	}

	for _, t := range all(data[1:]) {
		f := false
		pred := make([]byte, 0, len(data))
		pred = append(pred, data[0])

		s := string(data[0])
		for i, d := range t {
			s += string(d)
			if d == '(' {
				pred = append(pred, d)
				continue
			}
			if len(pred) == 0 || pred[len(pred)-1] != sim[d] {
				break
			}
			pred = pred[:len(pred)-1]
			if i == len(t)-1 && len(pred) == 0 {
				f = true
			}
		}

		if f {
			res = append(res, s)
		}
	}
	return res
}

func all(d []byte) [][]byte {
	replay := make(map[string]struct{})
	res := make([][]byte, 0)
	l := len(d)
	if l < 2 {
		return append(res, d)
	}

	switch l {
	case 3:
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				for k := 0; k < l; k++ {
					if i == k || j == k {
						continue
					}
					r := []byte{d[i], d[j], d[k]}
					if _, ok := replay[string(r)]; ok {
						continue
					}
					replay[string(r)] = struct{}{}
					res = append(res, r)
				}
			}
		}

	case 5:
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				for k := 0; k < l; k++ {
					if i == k || j == k {
						continue
					}
					for t := 0; t < l; t++ {
						if i == t || j == t || k == t {
							continue
						}
						for m := 0; m < l; m++ {
							if i == m || j == m || k == m || t == m {
								continue
							}
							r := []byte{d[i], d[j], d[k], d[t], d[m]}
							if _, ok := replay[string(r)]; ok {
								continue
							}
							replay[string(r)] = struct{}{}
							res = append(res, r)
						}
					}
				}
			}
		}

	case 7:
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				for k := 0; k < l; k++ {
					if i == k || j == k {
						continue
					}
					for t := 0; t < l; t++ {
						if i == t || j == t || k == t {
							continue
						}
						for m := 0; m < l; m++ {
							if i == m || j == m || k == m || t == m {
								continue
							}
							for n := 0; n < l; n++ {
								if i == n || j == n || k == n || t == n || m == n {
									continue
								}
								for p := 0; p < l; p++ {
									if i == p || j == p || k == p || t == p || m == p || n == p {
										continue
									}
									r := []byte{d[i], d[j], d[k], d[t], d[m], d[n], d[p]}
									if _, ok := replay[string(r)]; ok {
										continue
									}
									replay[string(r)] = struct{}{}
									res = append(res, r)
								}
							}
						}
					}
				}
			}
		}

	case 9:
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				for k := 0; k < l; k++ {
					if i == k || j == k {
						continue
					}
					for t := 0; t < l; t++ {
						if i == t || j == t || k == t {
							continue
						}
						for m := 0; m < l; m++ {
							if i == m || j == m || k == m || t == m {
								continue
							}
							for n := 0; n < l; n++ {
								if i == n || j == n || k == n || t == n || m == n {
									continue
								}
								for p := 0; p < l; p++ {
									if i == p || j == p || k == p || t == p || m == p || n == p {
										continue
									}
									for x := 0; x < l; x++ {
										if i == x || j == x || k == x || t == x || m == x || n == x || p == x {
											continue
										}
										for y := 0; y < l; y++ {
											if i == y || j == y || k == y || t == y || m == y || n == y || p == y || x == y {
												continue
											}
											r := []byte{d[i], d[j], d[k], d[t], d[m], d[n], d[p], d[x], d[y]}
											if _, ok := replay[string(r)]; ok {
												continue
											}
											replay[string(r)] = struct{}{}
											res = append(res, r)
										}
									}
								}
							}
						}
					}
				}
			}
		}

	case 11:
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				for k := 0; k < l; k++ {
					if i == k || j == k {
						continue
					}
					for t := 0; t < l; t++ {
						if i == t || j == t || k == t {
							continue
						}
						for m := 0; m < l; m++ {
							if i == m || j == m || k == m || t == m {
								continue
							}
							for n := 0; n < l; n++ {
								if i == n || j == n || k == n || t == n || m == n {
									continue
								}
								for p := 0; p < l; p++ {
									if i == p || j == p || k == p || t == p || m == p || n == p {
										continue
									}
									for x := 0; x < l; x++ {
										if i == x || j == x || k == x || t == x || m == x || n == x || p == x {
											continue
										}
										for y := 0; y < l; y++ {
											if i == y || j == y || k == y || t == y || m == y || n == y || p == y || x == y {
												continue
											}
											for a := 0; a < l; a++ {
												if i == a || j == a || k == a || t == a || m == a || n == a || p == a || x == a || y == a {
													continue
												}
												for b := 0; b < l; b++ {
													if i == b || j == b || k == b || t == b || m == b || n == b || p == b || x == b || y == b || a == b {
														continue
													}
													r := []byte{d[i], d[j], d[k], d[t], d[m], d[n], d[p], d[x], d[y], d[a], d[b]}
													if _, ok := replay[string(r)]; ok {
														continue
													}
													replay[string(r)] = struct{}{}
													res = append(res, r)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

	}
	return res
}

func noReplay(d [][]byte) [][]byte {
	b := make(map[string]struct{})
	res := make([][]byte, 0)

	for _, r := range d {
		s := string(r)
		_, ok := b[s]
		if !ok {
			b[s] = struct{}{}
			res = append(res, r)
		}
	}

	return res
}
