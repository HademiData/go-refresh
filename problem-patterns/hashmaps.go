package problempatterns

func TwoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {

		needed := target - num

		if idx, ok := seen[needed]; ok {

			return []int{i, idx}
		}
		seen[num] = i
	}

	return []int{}
}

func Anagram(a, b string) bool {

	if len(a) != len(b) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range a {

		count[ch]++

	}

	for _, ch := range b {

		count[ch]--

	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true

}

func firstNonRepeating(s string) int {

	count := make(map[rune]int)

	for _, ch := range s {

		count[ch]++
	}
	// for key,val := range count {

	// 	if val ==1 {
	// 		return key
	// 	}
	// }

	for i := 0; i < len(s); i++ {

		if count[rune(s[i])] == 1 {
			return i
		}
	}
	return 0
}

func IndexOfFirstOccurCharWithDuplicate(s string ) int{

	seen := make(map[byte]bool)

	for i:=0; i< len(s); i++ {

		if seen[s[i]] {
			return i
		}

		seen[s[i]] = true
	}
	return -1

}

func mostOccuringChar(s string ) rune {
	count := make(map[rune]int)

	for _,char := range s {
		count[char]++
	}
	max := count[rune(s[0])]
	var r rune

	for ch, cnt := range count {

		if cnt > max {
		
			r = ch
		}
	}
	return  r
}
