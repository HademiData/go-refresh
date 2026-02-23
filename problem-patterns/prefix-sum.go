package problempatterns

import "fmt"

// A prefix sum is where we pre-compute the cummulative sums of an array so we can answer range sum question very fast.

func prefixsum() {
	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}

	store := make([]int, len(nums))
	store = append(store, nums[0])

	for i := 1; i < (len(nums)); i++ {
		store = append(store, nums[i]+store[i-1])
	}
	fmt.Println(store)

}
