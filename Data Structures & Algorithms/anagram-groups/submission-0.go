func groupAnagrams(strs []string) [][]string {

    groups := make(map[[26]int][]string)

	for _, s := range strs{
		var key [26]int
		for _, char := range s{
		    key[char-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0 ,len(groups))

	for _, g := range groups{
		result = append(result, g)
	}

	return result
}
