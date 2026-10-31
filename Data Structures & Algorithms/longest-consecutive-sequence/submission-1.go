func longestConsecutive(nums []int) int {
    numSet := make(map[int]bool)

	max := 0
	
	for _, i := range nums{
		numSet[i] = true
	}

	for num := range numSet{
		if !numSet[num - 1]{
			tmp := 1
			for numSet[num + 1]{
				tmp++
				num++
			}
			if tmp >= max{
				max = tmp
			}
		}
	}

	return max

}
