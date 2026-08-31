func groupAnagrams(strs []string) [][]string {
	tr := make(map[string][]string)
	res := [][]string{}
	for _, value := range strs {
		runes := []rune(value)
		sort.Slice(runes, func(i, j int) bool{
			return runes[i] < runes[j]
		})
		tr[string(runes)] = append(tr[string(runes)], value)
	}
	for _, value := range tr {
		res = append(res, value)
	}
	return res
}

