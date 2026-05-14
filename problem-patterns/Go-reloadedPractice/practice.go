package goreloadedpractice

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
