package internal

func patternCount(text string, pattern string) int {
	count := 0
	T, P := len(text), len(pattern)
	for i := 0; i < T-P+1; i++ {
		if pattern == text[i:i+P] {
			count++
		}
	}
	return count
}

func frequentPattern(text string, k int) []string {
	maxFreq := 0
	patternFreq := make(map[string]int)
	N := len(text)
	for i := 0; i < N-k+1; i++ {
		patternFreq[text[i:i+k]]++
		if patternFreq[text[i:i+k]] > maxFreq {
			maxFreq = patternFreq[text[i:i+k]]
		}
	}
	maxList := make([]string, 0)
	for pattern, count := range patternFreq {
		if count == maxFreq {
			maxList = append(maxList, pattern)
		}
	}
	return maxList
}
