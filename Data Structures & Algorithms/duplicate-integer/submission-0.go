func hasDuplicate(nums []int) bool {
    duplicate := make(map[int]int, len(nums))
	for _,  num := range nums{
		if _, ok := duplicate[num]; ok{
			return true
		}
		duplicate[num] = 0
	}
	return false
}
