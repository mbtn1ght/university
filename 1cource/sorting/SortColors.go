package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	n := len(nums)
	for rep := 0; rep < len(nums)-1; rep++ {
		for i := 0; i < n-1; i++ {
			if nums[i] <= nums[i+1] {
				continue
			} else {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}

	}
	fmt.Print(nums)
}
