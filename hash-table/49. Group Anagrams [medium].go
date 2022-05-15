package hash_table

//思路:將每個字串轉為group(character count map), 然後去比對與現有group那些相同, 若相同就append 到特定index group的結果;
//若與現有任何group都不同, 即新增group, 以及append到新group的結果
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}

	groups := []map[byte]int{}
	for _, str := range strs {
		isfindgroup := false
		group2 := toGroup(str)
		for i, group := range groups {
			issame := isSameGroup(group, group2)
			if issame {
				result[i] = append(result[i], str)
				isfindgroup = true
				break
			}
		}

		if !isfindgroup {
			groups = append(groups, group2)
			result = append(result, []string{str})
		}

	}

	return result
}

func isSameGroup(group map[byte]int, group2 map[byte]int) bool {
	if len(group) != len(group2) {
		return false
	}

	for k, v := range group {
		if _, ok := group2[k]; !ok {
			return false
		}

		if group2[k] != v {
			return false
		}
	}

	return true
}

func toGroup(str string) map[byte]int {
	m := map[byte]int{}
	for _, b := range []byte(str) {
		if _, ok := m[b]; !ok {
			m[b] = 0
		}
		m[b]++
	}
	return m

}
