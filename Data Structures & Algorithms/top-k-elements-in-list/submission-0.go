func topKFrequent(nums []int, k int) []int {
    hash := make(map[int]int)

	for _,num := range nums{
		hash[num]++
	}

	arr := make([][2]int, 0, len(hash))
    for num, cnt := range hash {
        arr = append(arr, [2]int{cnt, num})
    }

    sort.Slice(arr, func(i, j int) bool {
        return arr[i][0] > arr[j][0]
    })

    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = arr[i][1]
    }
    return result
}
