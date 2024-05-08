package main

func main() {
	nums := []int{1, 2, 3}
	println(wiggleMaxLength(nums))
}

func wiggleMaxLength(nums []int) int {
	res := 1
	reverse := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] && reverse != 1 {
			res++
			reverse = 1
		} else if nums[i-1] > nums[i] && reverse != 2 {
			res++
			reverse = 2
		}
	}
	return res
}
