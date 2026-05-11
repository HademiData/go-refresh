package problempatterns

import "fmt"

// the two pointer pattern is used when we traverse through a list using two indices
// to move through an array or string instead of just one.
// instead of checking everything using nested loops
// we move two pointers in a smart way.



// Find two numbers in an ascending sorted array that sum up to a target
// e.g slice = [1,2,3,4,5] traget = 5

// ans 1, 4

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
}
