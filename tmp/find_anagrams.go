package tmp

func findAnagrams(s string, p string) []int {
	var pCounter, slideCounter [26]int
	ans := make([]int, 0)

	if len(p) > len(s) {
		return nil
	}

	for _, ch := range p {
		pCounter[ch-'a']++
	}
	for i := range p {
		slideCounter[s[i]-'a']++
	}

	for i := 0; ; i++ {
		if pCounter == slideCounter {
			ans = append(ans, i)
		}
		if i == len(s)-len(p) {
			break
		}
		slideCounter[s[i]-'a']--
		slideCounter[s[i+len(p)]-'a']++
	}

	return ans
}
