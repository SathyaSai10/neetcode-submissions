func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int, len(nums))

	for _, num := range nums{
		count[num]++
	}
 
    keys := make([]int, 0, len(count))
    for key := range count {
        keys = append(keys, key)
    }

    sort.Slice(keys, func(a, b int) bool {
        return count[keys[a]] > count[keys[b]]
    })

    return keys[:k]
}
