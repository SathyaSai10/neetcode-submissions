func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int, len(nums))
    for _, num := range nums {
        count[num]++
    }

    keys := make([][]int, len(nums)+1)
    for num, freq := range count {
        keys[freq] = append(keys[freq], num)
    }

    res := make([]int, 0, k)
    for f := len(keys) - 1; f > 0 && len(res) < k; f-- {
        for _, num := range keys[f] {
            res = append(res, num)
            if len(res) == k {
                return res
            }
        }
    }
    return res
}