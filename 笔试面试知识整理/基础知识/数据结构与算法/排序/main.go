package main

import "fmt"

func main() {
	fmt.Println("Sort")
	nums := []int{4, 3, 2, 1}
	// bubbleSort(nums)
	cocktailSort(nums)
	fmt.Println(nums)
}
