package main

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	println(canJump(nums))
}

func canJump(nums []int) bool {
	cover := 0
	for i := 0; i < cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums) {
			return true
		}
	}
	return false
}
