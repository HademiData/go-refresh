package goreloadedpractice

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Middle(words []string) []string {
	if len(words)%2 == 0 {

		return []string{"Even"}
	}
	for i := 0; i < len(words); i++ {
		if i == (len(words)-1)/2 {
			words[i] = strings.ToUpper(words[i])
		}
	}

	return words
}

func title(word string) string {
	return strings.Title(word)
}

func toUpper(word string) string {
	return strings.ToUpper(word)
}

func toLower(word string) string {
	return strings.ToLower(word)
}

func convHex(word string) (int64, error) {

	return strconv.ParseInt(word, 16, 64)
	//return val, nil
}

func convBin(word string) (int64, error) {

	return strconv.ParseInt(word, 2, 64)
	//return val, nil
}

func FixQuote(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func Article(word string) string {

	vowels := "aeiouh"
	VOWELS := "AEIOUH"
	if strings.Contains(vowels, string(word[0])) {
		return "an " + word
	}

	if strings.Contains(VOWELS, string(word[0])) {
		return "AN " + word
	}

	return "a " + word
}

func capLastN(words []string, N int) []string {

	start := len(words) - N
	if start < 0 {
		start = 0
	}

	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

func revStr(inpt string) string {

	text := []rune(inpt)
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {

		text[i], text[j] = text[j], text[i]
	}
	return string(text)
}

func save(num string, n int) string {
	if n < 0 {
		return num
	}

	ans := ""
	save := true
	for i := 0; i < len(num); i = i + n {

		end := i + n
		if end > len(num) {
			end = len(num)
		}

		if save {
			ans += num[i:end]
		}
		save = !save
	}
	return ans
}

func praSave(num string, n int) string {
	if len(num) == 1 || n == 0 {
		return num
	}
	save := true
	ans := ""
	for i := 0; i < len(num); i = i + n {

		end := i + n
		if end > len(num) {
			end = len(num)
		}

		if save {
			ans += num[i:end]
		}
		save = !save
	}
	return ans
}

func summ(n int) int {

	if len(strconv.Itoa(n)) < 2 {
		return n
	}

	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum

}

func ZipString(s string) string {

	if s == "" {
		return s
	}

	count := 1
	res := ""
	for i := 0; i < len(s); i++ {
		// if next char is same, count++
		if i+1 < len(s) && s[i] == s[i+1] {
			count++
		} else {
			// add current count and character
			res += fmt.Sprintf("%d%c", count, s[i])
			count = 1 // reset for next
		}
	}
	return res
}

func pracZip(s string) string {

	if s == "" {
		return s
	}

	res := ""
	count := 1
	for i := 0; i < len(s); i++ {

		if i+1 < len(s) && s[i] == s[i+1] {
			count++
		} else {
			res += fmt.Sprintf("%d%c", count, s[i])
			count = 1
		}
	}
	return res
}

func compare(str1 string, str2 string) string {
	st1 := []rune(str1)
	st2 := []rune(str2)

	for _, ch := range st1 {
		found := false
		for _, c := range st2 {

			if ch == c {
				found = true
				break
			}
		}
		if !found {
			return ""
		}
	}
	return str1
}

func isVow(rn rune) bool {
	r := unicode.ToLower(rn)
	return ('a' == r || 'e' == r || 'i' == r || 'o' == r || 'u' == r)
}

func countVow(s string) int {

	if s == "" {
		return 0
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if isVow(rune(s[i])) {
			count++
		}
	}
	return count
}

func isAlpha(r rune) bool {
	return ('A' <= r && 'Z' >= r) || 'a' <= r && 'z' >= r
}

func isPrint(r rune) bool {
	return 32 <= r && 126 >= r
}

func isNum(r rune) bool {
	return '0' <= r && '9' >= r
}

func wordCount(s string) int {

	sl := strings.Fields(s)
	m := make(map[string]int)

	if s == "" {
		return 0
	}

	for _, word := range sl {
		m[word]++
	}

	//delete(m, "g")

	return m["g"]
}

func recsumm(n int) int {

	if n == 0 {
		return 1
	}

	if len(strconv.Itoa(n)) < 2 {
		return n
	} else {
		return recsumm(recsumm(n/10) + n%10)
	}
}

func article(s string) string {

	if s == "" {
		return s
	}

	if strings.Contains("aeiouAEIOU", string(s[0])) {
		return "an " + s
	}

	if unicode.IsNumber(rune(s[0])) {
		return s
	}
	return "a " + s
}

func Chunk(slice []int, size int) [][]int {

	var res [][]int

	if size == 0 {
		fmt.Println()
		return res
	}

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		res = append(res, slice[i:end])
	}

	return res
}

func alpha(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func cipher(s string) string {

	out := ""
	for _, ch := range s {

		if 'a' <= ch && 'z' >= ch {
			m := 'z' - (ch - 'a')
			out += string(m)

		} else if 'A' <= ch && 'Z' >= ch {
			m := 'Z' - (ch - 'A')
			out += string(m)

		} else {
			out += string(ch)

		}
	}
	return out
}

func check(s, b string) string {

	if cipher(s) != b {
		return fmt.Sprintf("Err(CipherError { expected: %s })", cipher(s))
	}

	return "Ok(())"
}

func repAlph(s string) string {
	if s == "" {
		return s
	}

	out := ""

	for _, ch := range s {

		if 'A' <= ch && 'Z' >= ch {

			n := int(ch-'A') + 1
			for i := 0; i < n; i++ {
				out += string(ch)
			}
		} else if 'a' <= ch && 'z' >= ch {

			n := int(ch-'a') + 1
			for i := 0; i < n; i++ {
				out += string(ch)
			}

		} else {
			out += string(ch)
		}
	}

	return out
}

func alphaCount(s string) int {

	if s == "" {
		return 0

	}
	res := 0

	for _, ch := range s {
		if ('a' <= ch && 'z' >= ch) || ('A' <= ch && 'Z' >= ch) {
			res++
		}
	}
	return res
}

func AppendRange(min, max int) []int {

	if min >= max {
		return nil
	}

	var out []int

	for i := min; i < max; i++ {
		out = append(out, i)
	}

	return out
}

func HashCode(dec string) string {

	out := ""
	n := len(dec)

	for i := 0; i < n; i++ {
		val := (int(dec[i]) + n) % 127

		if val < 33 {
			val += 33
		}

		out += string(rune(val))
	}

	return out
}

func Capitalize(s string) string {

	if s == "" {
		return s
	}

	sl := strings.Fields(s)

	for i := 0; i < len(sl); i++ {
		sl[i] = strings.Title(sl[i])
	}

	return strings.Join(sl, " ")
}

func main() {

	words := []string{"hi", "yo", "gang", "cat", "fish"}

	fmt.Println(revStr("wale"))

	fmt.Println(Middle(words))

	fmt.Println(title("adewale"))
	fmt.Println(toLower("adeWALE"))
	fmt.Println(toUpper("adewale"))

	fmt.Println(convHex("1E"))
	fmt.Println(convBin("10"))
	fmt.Println(FixQuote("' hello '"))

	fmt.Println(Article("book"))
	fmt.Println(Article("house"))
	fmt.Println(Article("indomie"))

	//fmt.Println(capLastN(words, 1))

	re := regexp.MustCompile(`\s+`)
	input := "' awesome '"
	result := re.ReplaceAllString(input, "")
	fmt.Println(result)

}
