func getConcatenation(nums []int) []int {
    num := append(nums, nums...)
	return num
}
