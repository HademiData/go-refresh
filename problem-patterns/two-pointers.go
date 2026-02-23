package problempatterns

import "fmt"

// the two sum pattern is used when we traverse through a list using two indices
// to move through an array or string instead of just one.
// instead of checking everything using nested loops
// we move two pointers in a smart way.

func twopointers() {

	// both pointers in opposite directions
	nums := []int{1, 2, 3, 4, 6, 10}
	target := 8
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			fmt.Println("Found", nums[left], nums[right])
			break
		} else if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	// both pointers in same direction
	numbs := []int{112334455556677}
	slow := 0
	
	for fast := 1;fast < len(numbs); fast++ {
		if numbs[fast] != numbs[slow] {
			slow++
			numbs[slow] = numbs[fast]
		}

	}



}
