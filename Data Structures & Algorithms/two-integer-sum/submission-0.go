func twoSum(nums []int, target int) []int {
	store := make(map[int]int, len(nums))
    for idx, value := range nums{
		if idx1, ok := store[target - value]; ok{
            return []int{idx1,idx}
		}
        store[value] = idx
	}
	return []int{}
}
