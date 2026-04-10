func twoSum(nums []int, target int) []int {
	numHash := make(map[int]int)
    for i, num := range nums{
        diff := target - num
		if j, found := numHash[diff]; found{
			return []int{j, i}
		}
		numHash[num] = i
	}
	return []int{}
}
