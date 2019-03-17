package main

func insertionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ { // 假设第0个数已经有序，从第1个数开始，在有序序列中找到插入的位置
		for j := i - 1; j >= 0; j-- { // 在有序序列中从后往前查找插入的位置
			if nums[j] > nums[j+1] { // 通过交换来为这个数提供空位
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break // 不需要交换时，即找到了插入的位置，也无需继续比较了
			}
		}
	}
}
