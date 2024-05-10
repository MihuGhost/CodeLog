package main

func main() {
	nums := []int{2, 3, 1, 1, 4}
	println(jump(nums))
}

func jump(nums []int) int {
	end, maxPosition, res := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			res++
		}
	}
	return res
}
