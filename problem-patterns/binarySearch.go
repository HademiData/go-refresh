package problempatterns

func binarySearch(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	left := 0

	right := len(nums) - 1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1

		} else if nums[mid] > target {
			right = mid - 1

		} else if target == nums[mid] {

			return mid
		}
	}

	return -1
}

func SearchInsert(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	left := 0

	right := len(nums) - 1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1

		} else if nums[mid] > target {
			right = mid - 1

		} else if target == nums[mid] {

			return mid
		}
	}

	return left
}

func searchRange(nums []int, target int) []int {
	return []int{
		findFirst(nums, target),
		findLast(nums, target),
	}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}

		if nums[mid] == target {
			ans = mid
		}
	}

	return ans
}

func findLast(nums []int, target int) int {

	left, right := 0, len(nums)-1
	ans := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}

		if nums[mid] == target {
			ans = mid
		}
	}

	return ans
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {

			if target >= nums[left] &&
				target < nums[mid] {

				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {

			if target > nums[mid] &&
				target <= nums[right] {

				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// Koko eats bananas at speed k.

// Find MINIMUM speed needed to finish within h hours.

// Example:

// piles = [3,6,7,11]
// h = 8
func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0

	for _, p := range piles {
		if p > right {
			right = p
		}
	}

	ans := right

	for left <= right {

		mid := left + (right-left)/2

		hours := 0

		for _, p := range piles {
			hours += (p + mid - 1) / mid
		}

		if hours <= h {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

// Find minimum ship capacity needed to transport packages within days.

func shipWithinDays(weights []int, days int) int {

	left := 0
	right := 0

	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	for left < right {

		mid := left + (right-left)/2

		neededDays := 1
		cur := 0

		for _, w := range weights {

			if cur+w > mid {
				neededDays++
				cur = 0
			}

			cur += w
		}

		if neededDays <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
