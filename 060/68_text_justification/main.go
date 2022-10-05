/* https://leetcode.com/problems/text-justification/
68. Выравнивание текста

Учитывая массив строк word и ширину MaxWidth, отформатируйте текст таким образом, чтобы каждая строка содержала ровно символы MaxWidth и была полностью выровнена (слева и справа).
Вы должны упаковать свои слова с помощью жадного подхода; то есть упакуйте столько слов, сколько сможете, в каждую строку.
При необходимости добавьте дополнительные пробелы ' ', чтобы каждая строка содержала ровно символы максимальной ширины.
Дополнительные пробелы между словами должны быть распределены как можно более равномерно. Если количество пробелов в строке
не распределяется равномерно между словами, пустым ячейкам слева будет присвоено больше пробелов, чем ячейкам справа.
Что касается последней строки текста, то она должна быть выровнена по левому краю, и между словами не должно быть вставлено никаких дополнительных пробелов.

Примечание:
Слово определяется как последовательность символов, состоящая только из символов без пробелов.
Длина каждого слова гарантированно должна быть больше 0 и не превышать максимальную ширину.
Входной массив words содержит по крайней мере одно слово.

Пример 1:
Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

Пример 2:
Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]

Пояснение: Обратите внимание, что последняя строка - "shall be    " вместо "shall     be", потому что последняя строка должна быть выровнена по левому краю вместо полного выравнивания.
Обратите внимание, что вторая строка также выровнена по левому краю, поскольку содержит только одно слово.

Пример 3:
Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]

Выполнено:
Время выполнения: 4 мс, быстрее, чем 16,80% онлайн-заявок на выравнивание текста.
Использование памяти: 2,3 МБ, менее 8,80% отправленных онлайн-материалов для выравнивания текста.

*/

package main

import (
	"fmt"
)

func main() {
	//words, maxWidth := []string{"This", "is", "an", "example", "of", "text", "justification."}, 16
	//words, maxWidth := []string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16
	words, maxWidth := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20
	//words, maxWidth := []string{"Ab", "cd", "ef", "gh", "of", "tx", "."}, 14
	//words, maxWidth := []string{"a"}, 8
	text := fullJustify(words, maxWidth)
	printText(text)
}

func printText(text []string) {
	for _, d := range text {
		fmt.Printf("\"%s\"\n", d)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var idx int
	var line string
	res := make([]string, 0)
	for idx < len(words) {
		line, idx = buildLine(words, idx, maxWidth)
		res = append(res, line)
	}
	return res
}

func buildLine(words []string, idxA, maxWidth int) (string, int) {
	var line string
	var sum, sumLenWords int

	// определяем индекс последнего слова которое может поместиться в строке (idxB)
	idxB := idxA
	for i := idxA; i < len(words); i++ {
		l := len(words[i])
		sum += l
		if sum > maxWidth {
			break
		}
		idxB = i
		sumLenWords += l
		if sum == maxWidth || sum+1 == maxWidth {
			break
		}
		sum++
	}

	if idxB == len(words)-1 { // если последняя строка
		for i := idxA; i <= idxB; i++ {
			line += words[i]
			if i != idxB {
				line += whitespaceRepeat(1)
			}
		}
		line += whitespaceRepeat(maxWidth - len(line))
		return line, idxB + 1
	}

	// если в строке помещается только одно слово
	if idxA == idxB {
		line += words[idxA] + whitespaceRepeat(maxWidth-sumLenWords)
		return line, idxB + 1
	}

	// если в строке больше одного слова
	countWhitespace := (maxWidth - sumLenWords) / (idxB - idxA)
	remains := (maxWidth - sumLenWords) % (idxB - idxA)
	if countWhitespace == 0 {
		countWhitespace = 1
	}

	for i := idxA; i <= idxB; i++ {
		line += words[i]
		if i != idxB {
			line += whitespaceRepeat(countWhitespace)
			if remains > 0 {
				line += whitespaceRepeat(1)
				remains--
			}
		}
	}
	return line, idxB + 1
}

// строка из n пробелов
func whitespaceRepeat(n int) string {
	const whitespace = " "
	if n < 0 {
		return whitespace
	}

	var res string
	for i := 0; i < n; i++ {
		res += whitespace
	}
	return res
}
