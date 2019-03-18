package main

func quicksort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := nums[0]
	head := 0
	tail := len(nums) - 1
	for i := 1; i <= tail; {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	quicksort(nums[:head])
	quicksort(nums[head+1:])
}
