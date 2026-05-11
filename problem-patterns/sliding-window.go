package problempatterns

import (
	"math"
)

// finding the maximum sum of subarray with size k
// fixed size
func maxSum(nums []int, k int) int {
	
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum

	for i := k; i <  len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]

		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// variable size window (more important)

// Find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	
	left := 0
	maxLen := 0
	seen := make(map[byte]bool)

	for right := 0; right < len(s); right++ {

		// to check duplicates
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}

		// keeping track of the seen elements
		seen[s[right]] = true

		// updating the maxlen variable
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	
	return maxLen

}

// smallest Subarray with sum >= target
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			// shrink
			sum -= nums[left]
			left++
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen

}

// At Most K Distinct Characters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	left := 0
	maxLen := 0
	count := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		count[s[right]]++

		for len(count) > k {
			count[s[left]]--
			if count[s[left]] == 0 {
				delete(count, s[left])
			}
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}

	}

	return maxLen
}
