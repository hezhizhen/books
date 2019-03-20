package main

func heapsort(nums []int) {
	l := len(nums)
	mid := l / 2
	for i := mid; i > -1; i-- {
		heap(nums, i, l-1)
	}
	for i := l - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heap(nums, 0, i-1)
	}
}

func heap(nums []int, i, end int) {
	// parent: floor((i-1)/2)
	// root: i
	// left: 2i+1
	// right:2i+2
	l := 2*i + 1
	if l > end {
		return
	}
	n := l
	r := 2*i + 2
	if r <= end && nums[r] > nums[l] {
		n = r
	}
	if nums[i] > nums[n] {
		return
	}
	nums[n], nums[i] = nums[i], nums[n]
	heap(nums, n, end)
}
