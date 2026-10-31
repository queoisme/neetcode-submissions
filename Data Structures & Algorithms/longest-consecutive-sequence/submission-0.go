func longestConsecutive(nums []int) int {
    sort.Slice(nums, func(i, j int) bool{
        return nums[i] < nums[j]
    })
    if len(nums) == 0{
        return 0
    }
    max, tmp := 1, 1
    
	//m := make(map[int]int)
	for i := 1; i < len(nums); i++{
	    if nums[i] - nums[i-1] == 1{
	        tmp++
	    }else if nums[i] - nums[i - 1] == 0{
	        continue   
	    }else{
	        if tmp >= max{
	            max = tmp
	            tmp = 1
	        }
	    }
	}
	if tmp >= max {
	    max = tmp
	}
	return max
}
