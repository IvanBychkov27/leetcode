/* https://leetcode.com/problems/substring-with-concatenation-of-all-words/

Вам дается строка s и массив строковых слов одинаковой длины.
Возвращает все начальные индексы подстроки(ов) в s, которая представляет собой конкатенацию каждого слова в словах ровно один раз,
в любом порядке и без каких-либо промежуточных символов.
Вы можете вернуть ответ в любом порядке.

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Пояснение: Подстроки, начинающиеся с индексов 0 и 9, - это "barfoo" и "foobar" соответственно.
Порядок вывода не имеет значения, возврат [9,0] тоже подойдет.


Time Limit Exceeded
"bcabbcaabbccacacbabccacaababcbb"
["c","b","a","c","a","a","a","b","c"]

*/

package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	//s, words := "barfoothefoobarman", []string{"foo", "bar"} // [0,9]
	//s, words := "wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"} // []
	//s, words := "barfoofoobarthefoobarman", []string{"bar", "foo", "the"} // [6,9,12]
	//s, words := "wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"} // [8]
	//s, words := "bcabbcaabbccacacbabccacaababcbb", []string{"c", "b", "a", "c", "a", "a", "a", "b", "c"} // [6,16,17,18,19,20]
	//s, words := "baaaacbcc", []string{"bc", "ac", "aa", "ba"} // [0]
	s, words := "pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbcjjivtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbjjkgcztkcbwagtcnrtqryuqixtzhaakjlurnumzyovawrcjiwabuwretmdamfkxrgqgcdgbrdbnugzecbgyxxdqmisaqcyjkqrntxqmdrczxbebemcblftxplafnyoxqimkhcykwamvdsxjezkpgdpvopddptdfbprjustquhlazkjfluxrzopqdstulybnqvyknrchbphcarknnhhovweaqawdyxsqsqahkepluypwrzjegqtdoxfgzdkydeoxvrfhxusrujnmjzqrrlxglcmkiykldbiasnhrjbjekystzilrwkzhontwmehrfsrzfaqrbbxncphbzuuxeteshyrveamjsfiaharkcqxefghgceeixkdgkuboupxnwhnfigpkwnqdvzlydpidcljmflbccarbiegsmweklwngvygbqpescpeichmfidgsjmkvkofvkuehsmkkbocgejoiqcnafvuokelwuqsgkyoekaroptuvekfvmtxtqshcwsztkrzwrpabqrrhnlerxjojemcxel",
		[]string{"dhvf", "sind", "ffsl", "yekr", "zwzq", "kpeo", "cila", "tfty", "modg", "ztjg", "ybty", "heqg", "cpwo", "gdcj", "lnle", "sefg", "vimw", "bxcb"} // [ 935 ]

	//s, words := "faoucbmnvdujheznbntaszqsxicczlagnqbrsnfptbycfapjnkjflbzilemkpotehwvblcsefqgnsxwgkhnvjpgutuhklcosylvjonqtfyiyyegimtfxilrdiantcdncpiofxgaegjcenkobguyqhsupsjkxnxbehrjgxjlespiiazjiviyeaswuegtrexxnogumrfskwcuwbfgynfdpzzmdfhwletbvbvjtcbfydbxhgdfuiilkhznpjmcnhdkjytecujbykafqdkmovaacbntkbwyjziuaeycyhfytfdllqybnabpbqlmujmdiwxkaxnzeuxzcdknvkqimtzojkcdtoiemhonedokanjcswpnihvvaxlljprdfzjrhzgnfwyfkjhchyssfppfmaqwrjbwjmnslwhqsfverejacbshshohjhdaqgwzsmtfkbycitjzasccvukpcywlhboyjkzdyvjiwhngcwicqkhnercgrzuzcizmuyptvadrhqymmgqeqxrwqwfivzjzjklfkbygbczlszzhfpnxpfwfdbpacazlklqxordrveepedwlvmjktfrwihwkjvvugntweyqzcupgynstzdfskwqsfmcpixlqmenrcfsezjlxzdsyiztswjkdymsgldwxhlqlqjqsudptikuqjpihyslwgderjdqsqhejswbmzqihcczorhvrbiqhouaxxqroxvxrragssozvqajhhgakrqrfltekkvajtzbkzhfsvepnfawoiwzznsgammmykphdoipqrukobzbizxyhuxjjsrjgexgomntbyktphdekchsdfqmbqxkkpvstsyjfleilqbdxgjhkqbnvsbwkzguodmjwkubxaljqomouvxelztjtwhdvqzltlwpxssusffjtrznowtavmlojstgisuefsvclozdteopwnckmwmjkzavstecoyitsewduvjpzzexnjkbhykrympsitwwtfpllnrfiaukzzjoivrrueisqxmysiulpmzazqfkqcwrbfilbzcxfmrwmdwelrsbfrdehjqznmsquabxcfuhtlfhqcmbvgeaxkggvxfilxyfabecgalbnrjdtxhodnqcxwisvpahmyomztqhveljvumotteyhuagskzozbxlclabgslcwylruzhnvnlejnkcxlswnpjrajsjefnadauxzbmwrzaamnclauhplrgocbxggkjmkdllgykzzkamzcxazhpkywycxxlfhuttzfhhfrhedjqfnqfmxwzxuxztxmzgischzjrecajhjbmwrtlqqknmjpgg",
	//	[]string{"mntbyktphdekchsdfqmbq", "zznsgammmykphdoipqruk", "hyssfppfmaqwrjbwjmnsl", "hkqbnvsbwkzguodmjwkub", "qsqhejswbmzqihcczorhv", "xaljqomouvxelztjtwhdv", "dptikuqjpihyslwgderjd", "sozvqajhhgakrqrfltekk", "lszzhfpnxpfwfdbpacazl", "tavmlojstgisuefsvcloz", "wjkdymsgldwxhlqlqjqsu", "dteopwnckmwmjkzavstec", "obzbizxyhuxjjsrjgexgo", "rwqwfivzjzjklfkbygbcz", "whngcwicqkhnercgrzuzc", "frwihwkjvvugntweyqzcu", "rbiqhouaxxqroxvxrrags", "qmenrcfsezjlxzdsyizts", "cvukpcywlhboyjkzdyvji", "xkkpvstsyjfleilqbdxgj", "ykrympsitwwtfpllnrfia", "daqgwzsmtfkbycitjzasc", "whqsfverejacbshshohjh", "oyitsewduvjpzzexnjkbh", "izmuyptvadrhqymmgqeqx", "klqxordrveepedwlvmjkt", "qzltlwpxssusffjtrznow", "pgynstzdfskwqsfmcpixl", "vajtzbkzhfsvepnfawoiw"} // [0]

	res := findSubstring(s, words)
	fmt.Println("res =", res)

	//fmt.Println(pr(combinations(4)))

}

func findSubstring(s string, words []string) []int {
	res := []int{}
	if s == "" || len(words) == 0 {
		return res
	}
	lWords := len(words) * len(words[0])

	timeInterval := time.Now()

	for i := 0; i < len(s)-lWords+1; i++ {
		okInclusion := inclusion(s[i:i+lWords], words)
		if !okInclusion {
			continue
		}

		ok := []bool{false}
		compare(ok, s[i:i+lWords], words, 0)
		if ok[0] {
			if search(res, i) {
				continue
			}
			res = append(res, i)
		}

		if time.Now().Sub(timeInterval).Seconds() > float64(10) {
			fmt.Println("timeout output...")
			break
		}
	}

	fmt.Println(time.Now().Sub(timeInterval))

	return res
}

// вхождение
func inclusion(s string, ws []string) bool {
	for _, w := range ws {
		idx := strings.Index(s, w)
		if idx == -1 {
			return false
		}
	}
	return true
}

// вхождение
func inclusion1(s string, ws []string) ([]string, bool) {
	wIdx := make(map[int]string)
	id := make([]int, 0, len(ws))
	res := make([]string, 0, len(ws))

	for i, w := range ws {
		idx := strings.Index(s, w)
		if i == -1 {
			return res, false
		}
		wIdx[idx] = w
		id = append(id, i)
	}

	sort.Ints(id)

	for _, idx := range id {
		res = append(res, ws[idx])
	}

	return res, true
}

func compare(ok []bool, s string, words []string, n int) {
	if n == len(words)-1 {
		if s == strings.Join(words, "") {
			ok[0] = true
			return
		}
		return
	}

	for i := n; i < len(words); i++ {
		words[n], words[i] = words[i], words[n]
		compare(ok, s, words, n+1)
		words[n], words[i] = words[i], words[n]
		if ok[0] {
			return
		}
	}
}

func search(res []int, i int) bool {
	for _, d := range res {
		if d == i {
			return true
		}
	}
	return false
}
