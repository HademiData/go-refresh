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
