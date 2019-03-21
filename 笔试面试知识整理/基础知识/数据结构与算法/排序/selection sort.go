package main

func selectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i                               // 记录下标
		for j := min + 1; j < len(nums); j++ { // 第i趟：从i到len都是未排序，找到最小，同i交换，此时从0到i已排序，从i+1到len未排序
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}
