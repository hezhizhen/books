package main

func shellSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	mid := len(nums) / 2
	for mid > 0 {
		for i := mid; i < len(nums); i++ {
			j := i
			for j >= mid && nums[j] < nums[j-mid] {
				nums[j], nums[j-mid] = nums[j-mid], nums[j]
				j -= mid
			}
		}
		mid /= 2
	}
}
